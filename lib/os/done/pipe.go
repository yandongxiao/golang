package main

import "os"
import "fmt"
import "time"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func input(in *os.File) {
	for {
		buf := make([]byte, 1)
		in.Read(buf)
		fmt.Println(string(buf))
	}
}

func main() {
	in, out, err := os.Pipe()
	errcheck(err)
	go input(in)
	out.Write([]byte("nihao"))
	time.Sleep(5e9)
}
