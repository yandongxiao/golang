package main

import "fmt"

func main() {

	// golang不支持三目运算符: ternary if
	if 8%4 == 0 {
		println("if语句的基本用法, 条件两边的括号可以省，大括号不能省")
	}

	// A statement can precede conditionals;
	// any variables declared in this statement are available in all branches
	num := -10
	if num := 1; num < 0 { // 在更小的structure level定义了一个新的变量，并隐藏了同名的上级变量
		fmt.Printf("%d is negtive\n", num)
	} else if num == 0 {
		fmt.Printf("%d is zero\n", num)
	} else {
		fmt.Printf("%d is positive\n", num)
	}
	print(num)

	// error: non-bool val (type int) used as if condition)
	//val := 10
	// if val {
	//		println("val")
	//}
}
