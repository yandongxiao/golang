package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var (
		conn net.Conn
		err  error
	)

	if conn, err = net.Dial("tcp", "localhost:50000"); err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("please input your name: ")
	clientName, _ := inputReader.ReadString('\n')
	clientName = strings.Trim(clientName, "\r\n")

	for {
		fmt.Printf("input you words: ")
		line, _ := inputReader.ReadString('\n')
		line = strings.Trim(line, "\r\n")
		if line == "Q" {
			return
		}
		if _, err := conn.Write([]byte(clientName + " says: " + line + "\n")); err != nil {
			panic(err.Error())
		}
	}
}
