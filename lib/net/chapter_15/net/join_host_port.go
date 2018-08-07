package main

import "fmt"
import "net"

func main() {
	str := net.JoinHostPort("dxyan", "9000")
	fmt.Println(str)

}
