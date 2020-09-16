package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1e9)

	go func() {
		for now := range ticker.C {
			fmt.Println(now)
		}
	}()

	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	ticker.Stop()
}
