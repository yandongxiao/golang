package main

import "fmt"
import "net"

func main() {
	fmt.Println(net.LookupAddr("111.13.101.208")) /* ping baidu.com */
	fmt.Println(net.LookupAddr("192.168.0.45"))
}
