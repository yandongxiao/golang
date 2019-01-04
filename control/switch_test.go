// 传统的 switch 有三种应用方式
package main

import (
	"fmt"
	"time"
)

func ExampleNormal() {
	// 方式一：比较变量i与case值，与传统的switch语法相似
	i := 2
	switch i {
	case 1:
		// 不需要指定break，这是golang的默认行为
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	// Output:
	// Two
}

func ExampleMultipleExpressions() {
	// 方式二：You can use commas to separate multiple expressions
	// in the same case statement.
	switch time.Now().Weekday() {
	// NOTE: case 后面可以跟随多个表达式，它们之间是或的关系
	//后面没有接任何case语句，表明do nothing. 与其它case语句毫无关系
	case time.Monday, time.Tuesday, time.Friday,
		time.Saturday, time.Thursday:
	default:
		fmt.Println("It's a weekday")
	}

	// Output:
	//
}

func ExampleStatement() {
	// 扩展
	switch num := 10; num {
	case 10:
		fmt.Println("equal")
	default:
		fmt.Println("not equal")
	}

	// Output:
	// equal
}

func ExampleLikeIf() {
	v := 10
	// 方式三：switch without an expression is an alternate
	// way to express if/else logic.
	switch {
	case v < 12: // 这时候case表达式返回true or false
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Output:
	// It's before noon
}

func Example5() {
	// 扩展：与方式三的语义一样
	// 这里的分号是必须的，否则整个表达式会当做val进行处理
	switch num := 10; {
	case 1 == 1: // 需要是表达式形式
		fmt.Println(num)
	default:
		fmt.Println("not equal")
	}

	// Output:
	// 10
}
