package main

import "fmt"

func main() {
	// array的类型由两部分组成: [5]int
	var arrAge = [5]int{18, 20, 15, 22, 16}            // literal-1
	var arrLazy = [...]int{5, 6, 7, 8, 22}             // literal-2
	var arrKeyValue = [10]string{3: "Chris", 4: "Ron"} // literal-3
	var arr4 = [...]string{3: "Chris", 4: "Ron"}       // literal-4
	// var arrLazy = []int{5, 6, 7, 8, 22}			// 只要[]内什么都没有，返回的类型就是slice
	fmt.Printf("%T, %T, %T, %T\n", arrAge, arrLazy, arrKeyValue, arr4)
}
