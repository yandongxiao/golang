package main

import (
	"context"
	"fmt"
)

func ExampleCC() {
	ctx, _ := context.WithCancel(context.Background())
	ch1 := write(ctx, 0)
	ch2 := write(ctx, 1)
	for x := range sort(ctx, ch1, ch2) {
		fmt.Println(x)
	}
}

func write(ctx context.Context, begin int) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- begin:
				begin += 2
			}

			if begin >= 20 {
				close(ch)
				break
			}
		}
	}()
	return ch
}

func sort(ctx context.Context, ch1 <-chan int, ch2 <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		var v1, v2 *int
		ch1Closed, ch2Closed := false, false
		originCh1, originCh2 := ch1, ch2
		for ch1 != nil || ch2 != nil {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case v, ok := <-ch1:
				if !ok {
					ch1Closed = true
					ch1 = nil
				} else {
					v1 = &v
					ch1 = nil
				}
			case v, ok := <-ch2:
				if !ok {
					ch2Closed = true
					ch2 = nil
				} else {
					v2 = &v
					ch2 = nil
				}
			}

			if ch1Closed && ch2Closed {
				break
			} else if ch1Closed {
				ch <- *v2
				ch2 = originCh2
			} else if ch2Closed {
				ch <- *v1
				ch1 = originCh1
			} else if v1 != nil && v2 != nil {
				if *v1 > *v2 {
					ch <- *v2
					v2 = nil
					ch2 = originCh2
				} else {
					ch <- *v1
					v1 = nil
					ch1 = originCh1
				}
			}
		}
		close(ch)
	}()
	return ch
}
