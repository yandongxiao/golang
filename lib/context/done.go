package main

import (
	"context"
	"fmt"
)

func main() {
	// Done may return nil if this context can never be canceled.
	ctx := context.Background()
	fmt.Println(ctx.Done())
	ctx = context.WithValue(ctx, 10, 100)
	fmt.Println(ctx.Done())

	// Successive calls to Done return the same value.
	// WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline expires
	// WithTimeout arranges for Done to be closed when the timeout elapses.
	// NOTE: 如果协程在超时之前完成，那么还是要调用cancel来清除资源，而cancel是可重复调用的.
	ctx, _ = context.WithCancel(ctx)
	fmt.Println(ctx.Done(), ctx.Done())
}
