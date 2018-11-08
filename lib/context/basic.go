// NOTE: see block.go for problems
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chData := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		counter := 0
		for {
			time.Sleep(100 * time.Millisecond)
			counter++
			chData <- counter

			// check whether to end
			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}(ctx)

	// consumer
	for {
		counter := <-chData
		fmt.Println(counter)
		if counter > 3 {
			// go tool vet basic.go
			// basic.go:13: the cancel function returned by context.WithCancel should be called, not discarded, to avoid a context leak
			cancel()
			break
		}
	}
}
