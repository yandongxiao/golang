package main

import "fmt"

func get() (int, error) {
	// fmt.Errorf可以直接返回一个error的实例
	return 0, fmt.Errorf("%s", "error")
}

func main() {
	fmt.Println(get())
}
