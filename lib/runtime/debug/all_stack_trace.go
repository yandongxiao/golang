package main

import (
	"os"
	"runtime/pprof"
	"time"
)

func ExampleB() {
	go aaa()
	m111()
}

func m111() {
	m222()
}

func m222() {
	m333()
}

func m333() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	time.Sleep(time.Second)
}

func aaa() {
	time.Sleep(time.Second)
}
