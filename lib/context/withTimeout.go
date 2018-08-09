package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Second):
		fmt.Println("work done")
	}
}
