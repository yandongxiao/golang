package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	sleep := time.Millisecond * 1050
	// the cancel function returned by context.WithDeadline should be called, not discarded, to avoid a context leak
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	// Deadline returns the time when work done, 如果到截止时间，任务尚未完成，则自动执行
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Deadline())

	ch := make(chan int)
	go func(ctx context.Context) {
		select {
		case <-time.NewTimer(sleep).C:
			ch <- 1
		case <-ctx.Done():
			// NOTE: 需要返回一个值，哪怕这个值是表示错误
			ch <- 0
		}
	}(ctx)

	<-ch
	fmt.Println(ctx.Err())
}
