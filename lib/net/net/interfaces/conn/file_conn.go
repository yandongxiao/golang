// FileConn returns a copy of the network connection corresponding to the open file f.
// 在linux系统中，一切皆文件.
// 假设：go是C的一个子进程，go继承了C的pipe的读端，但是C只会告诉go这个读端的文件描述符.
//       go此时可以通过NewFile + FileConn将该文件描述符包装成net.Conn来使用
// It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

var end = make(chan bool)

func main() {
	socks, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}
	syscall.CloseOnExec(socks[0])
	syscall.CloseOnExec(socks[1])
	shimEnd := os.NewFile(uintptr(socks[0]), "shim end")
	connFile := os.NewFile(uintptr(socks[1]), "our end")
	defer shimEnd.Close()
	defer connFile.Close()

	// 证明了os.Pipe 不可用
	// shimEnd, connFile, _ = os.Pipe()

	go func() {
		conn1, _ := net.FileConn(shimEnd)
		defer conn1.Close()
		data := make([]byte, 100)
		n, err := conn1.Read(data)
		fmt.Println(n, string(data[:n]), err)
		end <- true
	}()

	conn2, err := net.FileConn(connFile)
	defer conn2.Close()
	conn2.Write([]byte("helloworld"))
	<-end
}
