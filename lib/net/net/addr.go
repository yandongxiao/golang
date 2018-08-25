/**
 * NOTE: this is an interface
 * type Addr interface {
 *    Network() string // name of the network (for example, "tcp", "udp")
 *    String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
 * }
 * The two methods Network and String conventionally return strings that can be passed as the arguments to Dial
 * the exact form and meaning of the strings is up to the implementation
 * 实现net.Addr接口的类型包括：*net.IPAddr, *net.TCPAddr, *net.UDPAddr, *net.UnixAddr
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	var listener net.Listener
	var err error

	if listener, err = net.Listen("tcp", "localhost:50000"); err != nil {
		panic("Listen: " + err.Error())
	}
	defer listener.Close()

	fmt.Println(listener.Addr().Network())
	fmt.Println(listener.Addr().String())
	addr := listener.Addr()
	fmt.Printf("%T, %s, %s\n", addr, addr.Network(), addr.String())

	// *net.TCPAddr
	tcpAddr := addr.(*net.TCPAddr)
	fmt.Printf("ip:%s, port:%d\n", tcpAddr.IP, tcpAddr.Port)

	// *net.IPNet
	ifi, _ := net.InterfaceByName("lo0")
	addrs, _ := ifi.Addrs()
	addr = addrs[0]
	ipAddr := addr.(*net.IPNet)
	fmt.Printf("ip:%s, mask:%v\n", ipAddr.IP, ipAddr.Mask)
}
