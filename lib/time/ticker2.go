package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := time.Tick(1000 * time.Millisecond)
	for t := range limiter {
		fmt.Println(t)
	}
}
