package main

import (
	"fmt"
	"time"
)

func Example_timestamp() {
	timeT, err := time.Parse("2006-01-02", "2020-04-14")
	if err != nil {
		panic(err)
	}

	timestamp := fmt.Sprintf("%v", timeT.UnixNano()/int64(time.Millisecond))
	fmt.Println(timeT)
	fmt.Println(timestamp)

	// Output:
	// 2020-04-14 00:00:00 +0000 UTC
	// 1586822400000
}
