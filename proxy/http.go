package proxy

import (
	"bytes"
	"crypto/tls"
	"sockets-proxy/util/log"

	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const ConnectSuccess = "HTTP/1.1 200 Connection Established\r\n\r\n"
const ConnectFailed = "HTTP/1.1 502 Bad Gateway\r\n\r\n"
const SslFileHost = "zt.io"

type ProxyHttp struct {
	client   Client
	request  *http.Request
	response *http.Response

	target net.Conn
	ssl    bool
	port   string
}

func NewProxyHttp() *ProxyHttp {
	p := &ProxyHttp{}
	return p
}

// Handle tcp连接处理入口
func (i *ProxyHttp) Handle() {
	request, err := http.ReadRequest(i.client.reader)
	if err != nil {
		return
	}
	i.port = "-1"
	if hostname := strings.Split(request.Host, ":"); len(hostname) > 1 {
		i.port = hostname[len(hostname)-1]
	}
	i.request = request
	// 如果是connect方法则是https请求或者ws、wss请求
	if i.request.Method == http.MethodConnect {
		i.ssl = true
		return
	}
	// 否则是普通请求
	i.ssl = false
	i.handleRequest()
}

// 处理请求入口
func (i *ProxyHttp) handleRequest() {
	var err error
	if i.request.URL == nil {
		log.Log.Println("请求地址为空")
		return
	}

	body, _ := i.ReadBody(i.request.Body)
	i.request.Body = io.NopCloser(bytes.NewReader(body))
	i.client.proxy.OnHttpRequestEvent(i.request)
	i.request.Body = io.NopCloser(bytes.NewReader(body))
	// 处理正常请求,获取响应
	i.response, err = i.Transport(i.request)
	if i.response == nil {
		log.Log.Println("远程服务器无响应-1")
		return
	}
	defer func() {
		if i.response.Body != nil {
			i.response.Body.Close()
		}
	}()
	if err != nil {
		log.Log.Println("获取远程服务器响应失败：" + err.Error())
		return
	}
	body, _ = i.ReadBody(i.response.Body)
	i.response.Body = io.NopCloser(bytes.NewReader(body))
	i.client.proxy.OnHttpResponseEvent(i.response)
	i.response.Body = io.NopCloser(bytes.NewReader(body))
	defer func() {
		if i.response.Body != nil {
			i.response.Body.Close()
		}
	}()
	_ = i.response.Write(i.client.conn)
}

// 读取http请求体
func (i *ProxyHttp) ReadBody(reader io.Reader) ([]byte, error) {
	if reader == nil {
		return []byte{}, nil
	}
	body, err := io.ReadAll(reader)
	return body, err
}

// 移除请求头
func (i *ProxyHttp) RemoveHeader(header http.Header) {
	removeHeaders := []string{
		"Keep-Alive",
		"Transfer-Encoding",
		"TE",
		"Connection",
		"Trailer",
		"Upgrade",
		"Proxy-Authorization",
		"Proxy-Authenticate",
		"Connection",
		"Accept-Encoding",
	}
	for _, value := range removeHeaders {
		if v := header.Get(value); len(v) > 0 {
			if strings.EqualFold(value, "Connection") {
				for _, customerHeader := range strings.Split(value, ",") {
					header.Del(strings.Trim(customerHeader, " "))
				}
			}
			header.Del(value)
		}
	}
}

// http请求转发
func (i *ProxyHttp) Transport(request *http.Request) (*http.Response, error) {
	// 去除一些头部
	i.RemoveHeader(request.Header)
	response, err := (&http.Transport{
		DisableKeepAlives:     true,
		ResponseHeaderTimeout: 60 * time.Second,
		//DialContext:           i.DialContext(),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}).RoundTrip(request)
	if err != nil {
		return nil, err
	}
	// 去除一些头部
	i.RemoveHeader(response.Header)
	return response, err
}

// 设置请求头
func (i *ProxyHttp) SetRequest(request *http.Request) *http.Request {
	request.Header.Set("Connection", "false")
	request.URL.Host = request.Host
	request.URL.Scheme = "https"
	return request
}
