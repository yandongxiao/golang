package main

import "net"
import "fmt"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func show_interface() {
	interfaces, err := net.Interfaces()
	errcheck(err)
	for idx := range interfaces {
		fmt.Println(interfaces[idx])
		fmt.Printf("Index = %v\n", interfaces[idx].Index)
		fmt.Printf("name = %v\n", interfaces[idx].Name)
		fmt.Printf("MTU = %v\n", interfaces[idx].MTU)
		fmt.Println("HardwareAddr =", interfaces[idx].HardwareAddr)

		MulticastAddrs, _ := interfaces[idx].MulticastAddrs()
		for i := range MulticastAddrs {
			fmt.Printf("MulticastAddrs[%d] = %v\n", i, MulticastAddrs[i])
		}

		Addrs, _ := interfaces[idx].Addrs()
		for i := range Addrs {
			fmt.Printf("Addrs[i] = %v\n", Addrs[i])
		}

		fmt.Printf("flags = %v\n", interfaces[idx].Flags)
	}
}

func main() {
	show_interface()
}
