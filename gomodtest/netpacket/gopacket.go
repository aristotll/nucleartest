package netpacket

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
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
