package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Info("info")
	glog.Warning("warn")
	glog.Error("error")

	fmt.Println(glog.Stats.Info.Lines(), glog.Stats.Info.Bytes())
	fmt.Println(glog.Stats.Warning.Lines(), glog.Stats.Warning.Bytes())
	fmt.Println(glog.Stats.Error.Lines(), glog.Stats.Error.Bytes())
}
