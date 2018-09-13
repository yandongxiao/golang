package main

import "fmt"

type addf func(a, b int) int

func add(a, b int) int {
	if a == 0 {
		panic("a==0")
	}
	return a + b
}

func foo(add addf) addf {
	return func(a, b int) (c int) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("---", err)
				c = -1
			}
		}()
		return add(a, b)
	}
}

func main() {
	add := foo(add)
	fmt.Println(add(0, 2))
}
