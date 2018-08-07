package main

// interface{} 不能省略
func test(nums ...interface{}) {
	for i, v := range nums {
		println(i, v)
	}
}

func main() {
	test(1.0, 2.0, 3.0)
	test(1, 2, 3)
}
