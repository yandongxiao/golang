package basic

import (
	"context"
	"fmt"
	"time"
)

func ExampleWithDeadline() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	ch := make(chan int)
	go func(ctx context.Context) {
		select {
		case <-time.NewTimer(time.Millisecond * 1050).C:
			ch <- 1
		case <-ctx.Done():
			ch <- 0
		}
	}(ctx)

	fmt.Println(<-ch)
	// Output:
	// 0
}
