package netpacket

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func SearchDevice() {
	// 得到所有的网络设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Devices found: ")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addressed: ", device.Addresses)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}

func SendPacket() {
	ipLayer := &layers.IPv4{
		SrcIP:    net.IP{127, 0, 0, 1},
		DstIP:    net.IP{127, 0, 0, 1},
		Protocol: layers.IPProtocolTCP,
	}
	// ethernetLayer := &layers.Ethernet{
	// 	SrcMAC: net.HardwareAddr{0xFF, 0xAA, 0xFA, 0xAA, 0xFF, 0xAA},
	// 	DstMAC: net.HardwareAddr{0xBD, 0xBD, 0xBD, 0xBD, 0xBD, 0xBD},
	// }
	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(8080),
	}
	if err := tcpLayer.SetNetworkLayerForChecksum(ipLayer); err != nil {
		log.Fatalln(err)
	}

	buffer := gopacket.NewSerializeBuffer()
	option := gopacket.SerializeOptions{}
	payload := gopacket.Payload([]byte{97, 98, 99})
	if err := gopacket.SerializeLayers(buffer, option, ipLayer, tcpLayer, payload); err != nil {
		log.Fatalln(err)
	}
	outgoingPacket := buffer.Bytes()

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write(outgoingPacket)
	if err != nil {
		log.Fatalln(err)
	}

	// fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer syscall.Close(fd)
	// addr := syscall.SockaddrInet4{
	// 	Port: 8080,
	// 	Addr: [4]byte{127, 0, 0, 1},
	// }
	// if err := syscall.Sendto(fd, outgoingPacket, 0, &addr); err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println("send success")
}

func SendRST(seq uint32) {
	ipLayer := &layers.IPv4{
		SrcIP:    net.IP{127, 0, 0, 1},
		DstIP:    net.IP{127, 0, 0, 1},
		Protocol: layers.IPProtocolTCP,
	}

	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(8080),
		RST:     true,
		Seq:     seq,
	}
	if err := tcpLayer.SetNetworkLayerForChecksum(ipLayer); err != nil {
		log.Fatalln(err)
	}

	buffer := gopacket.NewSerializeBuffer()
	option := gopacket.SerializeOptions{}
	if err := gopacket.SerializeLayers(buffer, option, ipLayer, tcpLayer); err != nil {
		log.Fatalln(err)
	}
	outgoingPacket := buffer.Bytes()

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write(outgoingPacket)
	if err != nil {
		log.Fatalln(err)
	}
}

func SendSYN() (nextACK uint32) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 2. 抓包拿到 ACK
		handle, err := pcap.OpenLive("lo0", 4096, false, time.Second*30)
		if err != nil {
			log.Fatalln(err)
		}
		defer handle.Close()

		source := gopacket.NewPacketSource(handle, handle.LinkType())
		packet := <-source.Packets()
		layer := packet.Layer(layers.LayerTypeTCP)
		t := layer.(*layers.TCP)
		//fmt.Printf("%+v\n", t)
		nextACK = t.Ack
		return
		// for packet := range source.Packets() {
		// 	layer := packet.Layer(layers.LayerTypeTCP)
		// 	t := layer.(*layers.TCP)
		// 	fmt.Printf("%+v\n", t)
		// 	nextACK = int64(t.Ack)
		// 	return
		// }
	}()

	// 1.先发送一个 SYN
	ipLayer := &layers.IPv4{
		SrcIP:    net.IP{127, 0, 0, 1},
		DstIP:    net.IP{127, 0, 0, 1},
		Protocol: layers.IPProtocolTCP,
	}

	tcpLayer := &layers.TCP{
		SrcPort: layers.TCPPort(4321),
		DstPort: layers.TCPPort(8080),
		SYN:     true,
	}
	if err := tcpLayer.SetNetworkLayerForChecksum(ipLayer); err != nil {
		log.Fatalln(err)
	}

	buffer := gopacket.NewSerializeBuffer()
	option := gopacket.SerializeOptions{}
	if err := gopacket.SerializeLayers(buffer, option, ipLayer, tcpLayer); err != nil {
		log.Fatalln(err)
	}
	outgoingPacket := buffer.Bytes()

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = conn.Write(outgoingPacket)
	if err != nil {
		log.Fatalln(err)
	}

	wg.Wait()
	return
}
