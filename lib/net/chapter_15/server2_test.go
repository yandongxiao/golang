package main

import (
	"fmt"
	"net"
	"testing"
)

func handleConn(t *testing.T, conn net.Conn) {
	buf := make([]byte, 3)
	n, err := conn.Read(buf)
	defer conn.Close()

	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("read a string:<%s>, len=%d", string(buf[:n]), n)
	n, err = conn.Read(buf)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("read a string:<%s>, len=%d", string(buf[:n]), n)

	n, err = conn.Write([]byte("world"))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("Write string len<%d>", n)
}

func Test1(t *testing.T) {
	ln, err := net.Listen("tcp", ":9527")
	if err != nil {
		t.Error(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			t.Error(err)
			return
		}

		go handleConn(t, conn)
	}
}
