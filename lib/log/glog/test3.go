package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func test() {
	glog.Infoln("helloworld")
}

func main() {
	flag.Parse()
	glog.V(2).Infoln("just info message")
	test()
	fmt.Println(glog.Stats.Info.Lines())

	fmt.Println(flag.Value())
	glog.Flush()
}
