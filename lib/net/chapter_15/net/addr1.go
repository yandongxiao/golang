package main

import "net"
import "fmt"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func show_avail() {
	/* structure addr has nothing except the address */
	addrs, err := net.InterfaceAddrs()
	errcheck(err)
	for idx := range addrs {
		fmt.Println(addrs[idx])
	}
}

func main() {
	show_avail()
}
