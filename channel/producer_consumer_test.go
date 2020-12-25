// NOTE：以下实现方案存在deadlock 的风险
// dead lock condition: producer wait on chData to send data;
// consumer wait on chEnd to send command to stop
package main

import (
	"fmt"
	"testing"
	"time"
)

func ExampleProducerConsumer() {
	chData := make(chan int)
	chEnd := make(chan struct{})

	// producer
	go func() {
		counter := 0
		for {
			// do work and send data
			// NOTE: 如果将该注释打开，则会出现 deadlock.
			time.Sleep(100 * time.Millisecond)
			counter++
			chData <- counter // blocked here

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
			chEnd <- struct{}{} // blocked here
			break
		}
	}
}

// solve problem
func TestProducerConsumer_2(t *testing.T) {
	chData := make(chan int)
	chEnd := make(chan struct{})

	// producer
	go func() {
		counter := 0
		for {
			time.Sleep(100 * time.Millisecond)
			counter++
			select {
			case <-chEnd:
				return
			case chData <- counter:
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
