package main

import "net"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	errcheck(err)

	for {
		_, err := ln.Accept() /* the new connection can not be closed by go */
		errcheck(err)
	}
}
