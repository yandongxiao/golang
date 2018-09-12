package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Error("error message")
	glog.Error("infomation message")
	glog.Flush()
}
