package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5e9)
	now := <-ticker.C
	fmt.Println(now)

}
