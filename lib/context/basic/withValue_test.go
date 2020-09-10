package basic

import (
	"context"
	"fmt"
)

func ExampleWithValue() {
	ctx := context.WithValue(context.Background(), 1, 100)
	fmt.Println(ctx.Value(1))

	ctx2 := context.WithValue(ctx, 1, 200)
	fmt.Println(ctx.Value(1))
	fmt.Println(ctx2.Value(1))

	// Output:
	// 100
	// 100
	// 200
}
