package icmp

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

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

// getHwAddr gets the hardware address of the gateway by sending an ARP request.
func (s *Scanner) getHwAddr() (net.HardwareAddr, error) {
	arpDst := s.gatewayAddr // arp 请求的目标地址，也就是网关地址

	// 定义一个以太网帧
	eth := &layers.Ethernet{}
	layers.ARP{}
}
