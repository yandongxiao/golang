package main

import "fmt"
import "net"
import "os"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

/*
 * fd, err := syscall.Socket
 * file := os.NewFile(uintptr(fd))
 * conn, err := net.FileConn(file)
 *
 */

/* The FileConn need a *os.File parameter, but not do like this*/
func main() {
	file, err := os.Open("/tmp/data")
	errcheck(err)

	conn, err := net.FileConn(file)
	errcheck(err)

	data := make([]byte, 1024)
	conn.Read(data)
	fmt.Println(string(data))

}
