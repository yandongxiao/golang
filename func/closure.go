package main

import "fmt"

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(300))
}

func Adder() func(int) int {
	num := 0
	return func(elm int) int {
		num += elm
		return num
	}
}
