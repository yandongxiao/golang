package main

import (
	"fmt"
	"time"
)

func main() {
	//for now := range time.Tick(time.Second) {
	//	fmt.Println(now)
	//}

        ticker := time.NewTicker(time.Second)
        for range ticker.C {
		fmt.Println("nn")
        }
}
