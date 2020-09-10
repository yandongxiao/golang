package basic

import (
	"context"
	"fmt"
	"time"
)

func ExampleWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 1秒钟后cancel会自动执行
	// 可重入、协程安全的函数
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("work done")
	}

	// Output:
	// context deadline exceeded
}
