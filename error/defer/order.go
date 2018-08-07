package main

import "fmt"

func guard() {
	if err := recover(); err != nil {
		fmt.Println("error message: ", err)
	}
}

func createpanic() {
	println("begin")
	defer guard() //NOTE the location
	defer println("defer createpanic")
	panic("try it")
	println("end")
}

func main() {
	createpanic()
	println("main end")
}
