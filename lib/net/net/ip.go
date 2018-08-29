package main

import (
	"fmt"
	"net"
)

func main() {
	ipaddr()
}

// 实现了net.Addr接口
func ipaddr() {
	// 方式一
	addr := &net.IPAddr{
		IP: net.ParseIP("192.168.12.1"),
	}
	fmt.Println(addr.Network(), addr.String())

	// 方式二
	addr, _ = net.ResolveIPAddr("ip", "192.168.12.1")
	fmt.Println(addr.Network(), addr.String())
}

func cidr() {
	ip, ipNet, err := net.ParseCIDR("192.0.2.1/30")
	if err != nil {
		panic(err)
	}
	fmt.Println("ip: ", ip, ip.DefaultMask())
	fmt.Println("ipNet: ", ipNet.IP, ipNet.Mask)
}

func basic() {
	// type IP []byte]
	// 虽然4个字节就可以存储一个IPv4的地址，但是len(ip)==16.
	// 所以，我们不能使用len(ip)来判断地址类型
	ip1 := net.IPv4(192, 168, 2, 4)
	fmt.Println(ip1, len(ip1))

	ip2 := net.ParseIP("192.168.2.4")
	fmt.Println(ip2, len(ip2))

	fmt.Println(ip1.Equal(ip2))

	// test
	ip := ip1
	fmt.Println("IsLoopback: ", ip.IsLoopback())
	fmt.Println("IsMulticast: ", ip.IsMulticast())
	fmt.Println("IsUnspecified: ", ip.IsUnspecified())
	fmt.Println("IsGlobalUnicast: ", ip.IsGlobalUnicast())
	fmt.Println("IsLinkLocalUnicast: ", ip.IsLinkLocalUnicast())
	fmt.Println("IsLinkLocalMulticast: ", ip.IsLinkLocalMulticast())
	fmt.Println("IsInterfaceLocalMulticast: ", ip.IsInterfaceLocalMulticast())

	// mask
	mask := net.IPMask([]byte{0xff, 0xff, 0xff, 0xf0})
	fmt.Println(ip.Mask(mask))

	fmt.Println(ip.MarshalText())
}
