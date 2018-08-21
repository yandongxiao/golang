package main

import (
	"fmt"
	"net"
)

func main() {
	// 既是network interface的名称和索引的映射关键，也是系统网络设备的信息
	// Interface represents a mapping between network interface name and index.
	// It also represents network interface facility information.
	interfaces, err := net.Interfaces()
	checkError(err)
	for idx := range interfaces {
		fmt.Printf("%+v\n", interfaces[idx])
	}

	// 通过network interface name or index来获取网络接口信息
	ifi, _ := net.InterfaceByName("lo0")
	fmt.Printf("%+v\n", ifi)
	ifi, _ = net.InterfaceByIndex(1)
	fmt.Printf("%+v\n", ifi)

	// Addrs returns a list of unicast interface addresses for a specific interface.
	// 网络接口设备和网络接口地址的关系是一对多的关系，即可以在一个接口上配置多个地址
	// e.g. 127.0.0.1/8, ::1/128, fe80::1/64
	addrs, _ := ifi.Addrs()
	for _, addr := range addrs {
		// ip+net 127.0.0.1/8
		// 这个地址的结果不能直接作为net.Dail的参数
		fmt.Printf("%T, %s, %s\n", addr, addr.Network(), addr.String())
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
