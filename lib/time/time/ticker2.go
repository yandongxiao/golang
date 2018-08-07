package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1e9)
	for now := range ticker.C {
		fmt.Println(now)
	}
}
