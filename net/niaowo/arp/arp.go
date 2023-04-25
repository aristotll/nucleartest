package main

import (
	"context"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/routing"
	"log"
	"net"
	"sync"
	"time"
)

// BroadcastAddr 是 MAC 广播地址
var BroadcastAddr = net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

type Scanner struct {
	iface               *net.Interface    // iface is the network interface on which to scan.
	gatewayAddr         net.IP            // gateway address.
	gatewayHardwareAddr *net.HardwareAddr //the gateway hardware address.
	srcIP               net.IP            // src is the source IP address.
	handle              *pcap.Handle      // handle is the pcap handle.
	// opts and buf allow us to easily serialize packets in the send() method.
	opts gopacket.SerializeOptions
	buf  gopacket.SerializeBuffer
}

func NewScanner() *Scanner {
	s := &Scanner{
		opts: gopacket.SerializeOptions{
			FixLengths:       true,
			ComputeChecksums: true,
		},
		buf: gopacket.NewSerializeBuffer(),
	}
	router, err := routing.New()
	if err != nil {
		log.Fatal(err)
	}
	// figure out the route by using the IP.
	iface, gw, src, err := router.Route(net.ParseIP("114.114.114.114"))
	if err != nil {
		log.Fatal(err)
	}
	s.gatewayAddr, s.srcIP, s.iface = gw, src, iface
	// open the handle for reading/writing.
	handle, err := pcap.OpenLive(iface.Name, 100, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	s.handle = handle
	gwHwAddr, err := s.getHwAddr()
	if err != nil {
		log.Fatal(err)
	}
	s.gatewayHardwareAddr = &gwHwAddr
	log.Printf("scanning with interface %v, gateway %v, src %v, hwaddr: %v", iface.Name, gw, src, gwHwAddr)
	return s
}

// getHwAddr gets the hardware address of the gateway by sending an ARP request.
func (s *Scanner) getHwAddr() (net.HardwareAddr, error) {
	arpDst := s.gatewayAddr // arp 请求的目标地址，也就是网关地址

	// 定义一个以太网帧
	eth := &layers.Ethernet{
		SrcMAC:       s.iface.HardwareAddr,
		DstMAC:       BroadcastAddr,
		EthernetType: layers.EthernetTypeARP,
	}
	// 构造一个 arp 请求报文，获取目标 IP 的 MAC 地址
	arp := &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,                            // MAC 地址的长度
		ProtAddressSize:   4,                            // IP 地址的长度
		Operation:         layers.ARPRequest,            // 操作类型，request or reply
		SourceHwAddress:   []byte(s.iface.HardwareAddr), // 源 MAC 地址
		SourceProtAddress: []byte(s.srcIP),              // 源 IP 地址
		DstHwAddress:      []byte{0, 0, 0, 0, 0, 0},     // 目的 MAC 地址，这里设置为广播地址
		DstProtAddress:    []byte(arpDst),               // 目的 IP 地址
	}
	// arp 报文需要封装在以太网帧中发送
	if err := s.sendPacket(eth, arp); err != nil {
		return nil, err
	}
	for {
		// 读取响应报文
		data, _, err := s.handle.ReadPacketData()
		if err != nil {
			return nil, err
		}
		// 超时未收到目标设备回应
		if err == pcap.NextErrorTimeoutExpired {
			log.Println("ARP request timed out")
			continue
		} else if err != nil {
			return nil, err
		}
		// 解析报文
		packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
		// 看响应的报文是不是 ARP 类型的
		if arpLayer := packet.Layer(layers.LayerTypeARP); arpLayer != nil {
			arp := arpLayer.(*layers.ARP)
			// 如果 ARP 响应报文的
			if net.IP(arp.SourceProtAddress).Equal(arpDst) {
				return arp.SourceHwAddress, nil
			}
		}
		return arp.SourceProtAddress, nil
	}
}

func (s *Scanner) sendPacket(l ...gopacket.SerializableLayer) error {
	if err := gopacket.SerializeLayers(s.buf, s.opts, l...); err != nil {
		return err
	}
	return s.handle.WritePacketData(s.buf.Bytes())
}

func (s *Scanner) send(input chan []string) error {
	return nil
}

func Arp(ctx context.Context, dstIP net.IP) {
	var wg sync.WaitGroup
	wg.Add(1)
	router, err := routing.New()
	if err != nil {
		log.Fatal(err)
	}
	// figure out the route by using the IP.
	iface, gw, srcIP, err := router.Route(net.ParseIP("114.114.114.114"))
	_ = gw
	if err != nil {
		log.Fatal(err)
	}
	// open the handle for reading/writing.
	handle, err := pcap.OpenLive(iface.Name, 100, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	// 定义一个以太网帧
	eth := &layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       BroadcastAddr, // 设置为广播地址
		EthernetType: layers.EthernetTypeARP,
	}
	// 构造一个 arp 请求报文，获取目标 IP 的 MAC 地址
	arp := &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,                                                    // MAC 地址的长度
		ProtAddressSize:   4,                                                    // IP 地址的长度
		Operation:         layers.ARPRequest,                                    // 操作类型，request or reply
		SourceHwAddress:   iface.HardwareAddr,                                   // 源 MAC 地址
		SourceProtAddress: srcIP,                                                // 源 IP 地址
		DstHwAddress:      net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // 目的 MAC 地址，这里设置为广播地址
		DstProtAddress:    dstIP.To4(),                                          // 目的 IP 地址，这里一定要调用 To4（不知道为啥）
	}

	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	buf := gopacket.NewSerializeBuffer()
	// 分开方便 debug
	if err := gopacket.SerializeLayers(buf, opts, eth, arp); err != nil {
		log.Fatalln(err)
	}
	//if err := gopacket.SerializeLayers(buf, opts, arp); err != nil {
	//	log.Fatalln(err)
	//}
	if err := handle.WritePacketData(buf.Bytes()); err != nil {
		log.Fatalln(err)
	}
	go func() {
		defer wg.Done()
		ps := gopacket.NewPacketSource(handle, handle.LinkType())
		// wait 3 seconds for an ARP reply.
		for {
			select {
			case <-ctx.Done():
				log.Println("wait reply timeout")
				return
			case p := <-ps.Packets():
				// 看响应的报文是不是 ARP 类型的
				if arpLayer := p.Layer(layers.LayerTypeARP); arpLayer != nil {
					arp := arpLayer.(*layers.ARP)
					if arp.Operation != layers.ARPReply {
						log.Println("not ARP reply")
						log.Println(arp.Operation, arp.DstProtAddress)
						continue
					}
					// 如果 ARP 响应报文的源 IP 地址
					if net.IP(arp.SourceProtAddress).Equal(dstIP) {
						fmt.Println(arp.SourceProtAddress)
					}
				}
			}
			// 读取响应报文
			//data, _, err := handle.ReadPacketData()
			//if err != nil {
			//	log.Fatalln(err)
			//}
			//// 超时未收到目标设备回应
			//if err == pcap.NextErrorTimeoutExpired {
			//	log.Println("ARP request timed out")
			//	continue
			//} else if err != nil {
			//	log.Fatalln(err)
			//}
			//// 解析报文
			//packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
			////println(packet.String())
			//// 看响应的报文是不是 ARP 类型的
			//if arpLayer := packet.Layer(layers.LayerTypeARP); arpLayer != nil {
			//	arp := arpLayer.(*layers.ARP)
			//	// 如果 ARP 响应报文的源 IP 地址
			//	if net.IP(arp.SourceProtAddress).Equal(dstIP) {
			//		fmt.Println(arp.SourceProtAddress)
			//	}
			//}
		}
	}()

	wg.Wait()
}

// apt install -y libpcap-dev
// apt install -y libpcap-dev
//
// Mac 下交叉编译：
// brew tap messense/macos-cross-toolchains
// # install x86_64-unknown-linux-gnu toolchain
// brew install x86_64-unknown-linux-gnu
// # install aarch64-unknown-linux-gnu toolchain
// brew install aarch64-unknown-linux-gnu
//
// 编译:
// GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-unknown-linux-gnu-gcc CGO_CFLAGS="-L/opt/homebrew/opt/libpcap/lib" go build -o arp_linux
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	Arp(ctx, net.ParseIP("192.168.31.80"))
}
