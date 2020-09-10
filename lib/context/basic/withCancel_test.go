package basic

import (
	"context"
	"fmt"
	"time"
)

func ExampleWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			case <-time.After(time.Second):
				fmt.Println("working")
			}
		}
	}(ctx)

	time.Sleep(2500 * time.Millisecond)
	cancel()
	// Output:
	// working
	// working
}
