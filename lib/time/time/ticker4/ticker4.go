package main

import (
	"fmt"
	"time"
)

func main() {
	for now := range time.Tick(time.Second) {
		fmt.Println(now)
	}
}
