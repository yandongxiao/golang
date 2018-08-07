package main

import (
	"time"
)

func main() {
	chRate := time.Tick(time.Second)
	for {
		<-chRate
		println("a message")
	}

}
