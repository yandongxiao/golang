package main

import (
	"fmt"
	"time"
)

func ExampleTimestamp() {
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano()/int64(time.Millisecond))
	fmt.Println(timestamp)
}
