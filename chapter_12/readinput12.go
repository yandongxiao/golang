package main

import "fmt"

func test1() {

	var name string
	var age int
	n, err := fmt.Scanln(&name, &age)
	if err == nil {
		fmt.Println("read", n, " data :", name, age)
	} else {
		fmt.Println("read error")
	}
}

func test2() {
	var name string
	var age int

	fmt.Sscanf("ydx 19", "%s %d", &name, &age)
	println(name, age)
}

func test3() {
	var name string
	var age int
	fmt.Scanf("%s %d", &name, &age)
	println(name, age)
}

func main() {
	test1()
}
