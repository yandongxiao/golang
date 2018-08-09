package main

import (
	"context"
	"fmt"
)

// should not use basic type int as key in context.WithValue
type INT int

func main() {
	ctx := context.WithValue(context.Background(), INT(1), 100)
	fmt.Println(ctx.Value(INT(1))) // 100

	// 继承了ctx, 但是每一个context都拥有自己独立的map
	ctx2 := context.WithValue(ctx, INT(1), 200)
	fmt.Println(ctx.Value(INT(1)))  // 100
	fmt.Println(ctx2.Value(INT(1))) // 200
}
