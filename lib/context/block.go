// NOTE：以下实现方案存在dead lock的风险
// dead lock condition: producer wait on chData to send data;
// consumer wait on chEnd to send command to stop
package main

import (
	"fmt"
	"time"
)

func main() {
	chData := make(chan int)
	chEnd := make(chan struct{})

	go func() { // producer
		counter := 0
		for {
			// do work and send data
			// NOTE: 如果将该注释打开，则会出现deal lock.
			time.Sleep(100 * time.Millisecond)
			counter++
			chData <- counter

			// check whether to end
			select {
			case <-chEnd:
				return
			default:
			}
		}
	}()

	// consumer
	for {
		counter := <-chData
		fmt.Println(counter)
		if counter > 3 {
			chEnd <- struct{}{}
			break
		}
	}
}
