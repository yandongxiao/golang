package main

import (
	"flag"
	"github.com/golang/glog"
)

func test() {

}

func main() {
	flag.Parse()
	test()
	glog.Error("infomation message")
	glog.Error("infomation message")
	glog.Flush()
}
