package main

func test1(a, b, c int) int {
	return a + b + c
}

func test2(a, b int) (int, int, int) {
	return a, b, 3
}

func test3(a, b int) (int, int) {
	return a, b
}

func main() {
	// test1(test3(1,2), 3) 这种调用方式是不被允许的
	println(test1(test2(1, 2)))
}
