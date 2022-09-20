package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

var CIDR = flag.String("ip", "127.0.0.0/8", "cidr ip")

func main() {
	flag.Parse()
	ip, ipnet, err := net.ParseCIDR(*CIDR)
	if err != nil {
		panic(err)
	}
	ones, bits := ipnet.Mask.Size()
	fmt.Printf("ip: %v, ipnet: %v, ipnet.Mask: %v, mask.size: %v/%v\n", ip, ipnet, ipnet.Mask, ones, bits)
	binaryMask := getMaskBinary(ones, bits)
    decimalMask := getMaskDecimal(binaryMask)
    fmt.Printf("%v => %v\n", binaryMask, decimalMask)
}

// 将子网掩码转换为二进制形式的表示，比如 255.128.0.0 => 11111111.10000000.00000000.00000000
func getMaskBinary(ones, bits int) string {
	count1 := ones
	count0 := bits - ones

	var has8 int8
	var ret strings.Builder
	for i := 0; i < count1; i++ {
		if has8 == 8 {
			ret.WriteString(".")
			has8 = 0
		}
		ret.WriteString("1")
		has8++
	}

	for i := 0; i < count0; i++ {
		if has8 == 8 {
			ret.WriteString(".")
			has8 = 0
		}
		ret.WriteString("0")
		has8++
	}

	return ret.String()
}

func getMaskDecimal(dec string) string {
	var sb strings.Builder
	var count int8
	s2 := strings.Split(dec, ".")

	for _, v := range s2 {
		i, err := strconv.ParseInt(v, 2, 64)
		if err != nil {
			panic(err)
		}
		sb.WriteString(strconv.Itoa(int(i)))
		if count < 3 {
			sb.WriteString(".")
            count++
		}
	}
	return sb.String()
}
