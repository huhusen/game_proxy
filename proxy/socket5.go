package proxy

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"sockets-proxy/log"
	"sockets-proxy/util"
	"strconv"
	"time"
)

type Socket5 struct {
	client Client
	port   string
	remote net.Conn
	state  bool
}

const (
	Rsv          = 0x00 //预留位
	CommandConn  = 0x01 // 命令
	CommandBind  = 0x02
	CommandUdp   = 0x03
	TargetIpv4   = 0x01 // 目标类型
	TargetIpv6   = 0x04
	TargetDomain = 0x03
	Version      = 0x5
	SocketServer = "server"
	SocketClient = "client"
)

func (s *Socket5) Send2Server(data []byte) {
	if s.state {
		s.remote.Write(data)
	}
}
func (s *Socket5) Send2client(data []byte) {
	if s.state {
		s.client.conn.Write(data)
	}
}
func (s *Socket5) Handle() {
	// 读取版本号
	version, err := s.client.reader.ReadByte()
	if err != nil {
		log.Log.Println("读取socket5版本号错误：" + err.Error())
		return
	}
	if version != Version {
		log.Log.Println("socket5版本号不匹配")
		return
	}
	// 读取支持的方法
	methodNum, err := s.client.reader.ReadByte()
	if err != nil {
		log.Log.Println("读取socket5支持方法数量错误：" + err.Error())
		return
	}
	if methodNum < 0 || methodNum > 0xFF {
		log.Log.Println("socket5支持方法参数错误")
		return
	}
	// 是否需要账号密码验证
	var requiredAuth bool
	method := uint8(0x00)
	// 读取所有的方法列表
	for n := 0; n < int(methodNum); n++ {
		method, err = s.client.reader.ReadByte()
		if err != nil {
			log.Log.Println("读取socket5支持错误：" + err.Error())
			return
		}
		if method == 0x02 {
			requiredAuth = true
		}
	}

	_, err = s.client.writer.Write([]byte{version, method})
	if err != nil {
		log.Log.Println("返回数据错误：" + err.Error())
		return
	}
	_ = s.client.writer.Flush()
	if requiredAuth {
		// TODO 账号密码验证
		return
	}
	// 读取版本号
	version, err = s.client.reader.ReadByte()
	if version != Version {
		log.Log.Println("socket5版本号错误")
		return
	}
	// 读取命令
	command, err := s.client.reader.ReadByte()
	if err != nil {
		log.Log.Println("读取socket5命令错误")
		return
	}
	if command != CommandConn && command != CommandBind && command != CommandUdp {
		log.Log.Println("不支持socket5命令")
		return
	}
	// 读取保留位
	rsv, err := s.client.reader.ReadByte()
	if err != nil || rsv != Rsv {
		log.Log.Println("读取socket5保留位错误")
		return
	}
	// 读取目标地址类型
	targetType, err := s.client.reader.ReadByte()
	if err != nil {
		log.Log.Println("读取socket5保留位错误")
		return
	}
	if targetType != TargetIpv4 && targetType != TargetIpv6 && targetType != TargetDomain {
		log.Log.Println("不支持socket5地址")
		return
	}
	var hostname string
	switch targetType {
	case TargetIpv4:
		buffer := make([]byte, 4)
		// 读4字节
		n, err := s.client.reader.Read(buffer)
		if err != nil || n != len(buffer) {
			log.Log.Println("读取ipv4地址错误")
			return
		}
		hostname = net.IP(buffer).String()
		break
	case TargetIpv6:
		buffer := make([]byte, 16)
		// 读16字节
		n, err := s.client.reader.Read(buffer)
		if err != nil || n != len(buffer) {
			log.Log.Println("读取ipv6地址错误")
			return
		}
		hostname = net.IP(buffer).String()
		break
	case TargetDomain:
		// 读取域名长度
		domainLen, err := s.client.reader.ReadByte()
		if err != nil || domainLen <= 0 {
			log.Log.Println("读取域名地址错误")
			return
		}
		buffer := make([]byte, domainLen)
		n, err := s.client.reader.Read(buffer)
		if err != nil || n != len(buffer) {
			log.Log.Println("读取域名地址错误")
			return
		}
		addr, err := net.ResolveIPAddr("ip", string(buffer))
		if err != nil {
			log.Log.Println("读取域名地址错误：" + err.Error())
			hostname = string(buffer)
		} else {
			hostname = addr.String()
		}
		break
	}
	// 读端口号,大端
	buffer := make([]byte, 2)
	_, err = s.client.reader.Read(buffer)
	if err != nil {
		log.Log.Println("读取端口号错误：" + err.Error())
		return
	}
	s.port = strconv.Itoa(int(util.ByteToInt(buffer)))
	hostname = fmt.Sprintf("%s:%s", hostname, s.port)
	// 写入版本号
	_ = s.client.writer.WriteByte(Version)
	if command == 0x03 {
		s.remote, err = net.DialTimeout("udp", hostname, time.Second*30)
	} else {
		if s.port == "443" {
			dialer := &net.Dialer{
				Timeout: time.Second * 30,
			}
			s.remote, err = tls.DialWithDialer(dialer, "tcp", hostname, &tls.Config{
				InsecureSkipVerify: true,
			})
		} else {
			s.remote, err = net.DialTimeout("tcp", hostname, time.Second*30)
		}
	}
	log.Log.Println("待连接的目标服务器：" + hostname)
	// 写入Rep
	if err != nil {
		log.Log.Println("连接目标服务器失败：" + hostname + " " + err.Error())
		_ = s.client.writer.WriteByte(0x01)
		_ = s.client.writer.Flush()
		return
	} else {
		_ = s.client.writer.WriteByte(0x00)
	}
	// 写入Rsv
	_ = s.client.writer.WriteByte(Rsv)
	remoteAddr := s.remote.RemoteAddr().String()
	host, _, _ := net.SplitHostPort(remoteAddr)
	if util.IsIpV4(host) {
		_ = s.client.writer.WriteByte(TargetIpv4)
		_, _ = s.client.writer.Write(net.ParseIP(host).To4())
	}
	if util.IsIpV6(host) {
		_ = s.client.writer.WriteByte(TargetIpv6)
		_, _ = s.client.writer.Write(net.ParseIP(host).To16())
	}
	if !util.IsIpV4(host) && !util.IsIpV6(host) {
		_ = s.client.writer.WriteByte(TargetDomain)
		_ = s.client.writer.WriteByte(byte(len(hostname)))
		_, _ = s.client.writer.WriteString(hostname)
	}
	// 写入端口
	_, _ = s.client.writer.Write(buffer)
	err = s.client.writer.Flush()
	if err != nil {
		log.Log.Println("写入socket5握手错误：" + err.Error())
		return
	}
	s.state = true
	out := make(chan error, 2)
	if command == 0x01 {
		go s.Transport(out, s.client.conn, s.remote, SocketClient)
		go s.Transport(out, s.remote, s.client.conn, SocketServer)
		<-out
	}
}

func (s *Socket5) Transport(out chan<- error, originConn net.Conn, targetConn net.Conn, role string) {
	buff := make([]byte, 10*1024)
	for {
		readLen, err := originConn.Read(buff)
		if readLen > 0 {
			buff = buff[0:readLen]
			if role == SocketServer {
				s.client.proxy.OnSocket5ResponseEvent(buff)
			} else {
				s.client.proxy.OnSocket5RequestEvent(buff)
			}
			writeLen, err := targetConn.Write(buff)
			if writeLen < 0 || readLen < writeLen {
				writeLen = 0
				if err == nil {
					out <- errors.New("写入目标服务器错误-1")
					break
				}
			}
			if readLen != writeLen {
				out <- errors.New("写入目标服务器错误-2")
				break
			}
		}
		if err != nil {
			if err != io.EOF {
				out <- errors.New("读取客户端数据错误-1")
			}
			break
		}
		buff = buff[:]
	}
}
