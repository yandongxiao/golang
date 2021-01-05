package main

import (
	"flag"
)

var (
	loglevel int
	debug    bool
)

func init() {
	flag.IntVar(&loglevel, "loglevel", 1, "the log level")
	flag.BoolVar(&debug, "debug", false, "the debug")
}

func main2() {
	flag.Usage()
}
