package pcap

import (
	"encoding/json"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"sockets-proxy/log"
)

type Net struct {
	Ethernet struct {
		SrcMAC       string
		DstMAC       string
		EthernetType string
	}
	IPV4 struct {
		SrcIP    string
		DstIP    string
		Protocol string
	}
	TCP struct {
		SrcPort string
		DstPort string
		Seq     uint32
	}
}

func (n Net) String() string {
	b, _ := json.Marshal(&n)
	return string(b)
}

type CaptureUnPack struct {
	file       string
	Filter     string
	DataHandle func(data []byte, n Net)
}

func NewCaptureUnPack(file string) *CaptureUnPack {
	return &CaptureUnPack{file: file}
}
func (c CaptureUnPack) Run() {
	handle, err := pcap.OpenOffline(c.file)
	if err != nil {
		log.Log.Errorln("read pcap file error.", err)
	}
	if c.Filter != "" {
		handle.SetBPFFilter(c.Filter)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		c.packetInfo(packet)
	}
	defer handle.Close()
}
func (c CaptureUnPack) packetInfo(packet gopacket.Packet) {
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	net := Net{}
	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		net.Ethernet.SrcMAC = ethernetPacket.SrcMAC.String()
		net.Ethernet.DstMAC = ethernetPacket.DstMAC.String()
		net.Ethernet.EthernetType = ethernetPacket.EthernetType.String()

	}
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		net.IPV4.SrcIP = ip.SrcIP.String()
		net.IPV4.DstIP = ip.DstIP.String()
		net.IPV4.Protocol = ip.Protocol.String()
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		net.TCP.SrcPort = tcp.SrcPort.String()
		net.TCP.DstPort = tcp.DstPort.String()
		net.TCP.Seq = tcp.Seq
	}

	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		if c.DataHandle != nil {
			c.DataHandle(applicationLayer.Payload(), net)
		}
	}

	if err := packet.ErrorLayer(); err != nil {
		log.Log.Errorln("Error decoding some part of the packet:", err)
	}
}
