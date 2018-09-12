package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)
		infof  = func(info string) {
			logger.Output(2, info)
		}
	)

	infof("Hello")
	fmt.Print(&buf)

	logger.SetOutput(os.Stdout)
	logger.Output(1, "world")
}

func main2() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	f := func(s string) {
		func(s string) {
			func(s string) {
				log.Output(4, s)
			}(s)
		}(s)
	}

	f("hello") // 2018/09/12 14:42:57 output2.go:15: hello
	f("world") // 2018/09/12 14:42:57 output2.go:16: world
}
