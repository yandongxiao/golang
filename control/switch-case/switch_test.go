/*
	switch InitSimpleStatement; CompareOperand0 {
	case CompareOperandList1:
		// do something
	case CompareOperandList2:
		// do something
	...
	case CompareOperandListN:
		// do something
	default:
		// do something
	}
*/
package main

import (
	"fmt"
	"time"
)

// 不需要指定break，这是golang的默认行为
func ExampleNormal() {
	//x := int64(3)
	type INT int
	switch 1 { // 2是无类型的常量，隐式地将它转换为int类型.
	case 1: // 1是无类型的常量，隐式地将它转换为与CompareOperand0为相同的类型
		fmt.Println("One")
		// invalid case INT(2) in switch on 2 (mismatched types INT and int)
		//case INT(2):
		//	fmt.Println("Two")
		// invalid case x in switch on 2 (mismatched types int64 and int)
		// case x:
		// fmt.Println("Three")
	}

	// Output:
	// One
}

// 方式二：You can use commas to separate multiple expressions
// in the same case statement.
func ExampleMultipleExpressions() {
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

// InitSimpleStatement
func ExampleStatement() {
	switch num := 10; num {
	case 10:
		fmt.Println("equal")
	default:
		fmt.Println("not equal")
	}

	// Output:
	// equal
}

func Example5() {
	// 这里的分号是必须的，否则编译器报错。
	switch num := 10; { // CompareOperand0的值为true，类型为bool
	case 1 == 1: // 表达式的结果也必须是bool类型的值，这样才能与CompareOperand0进行比较
		fmt.Println(num)
	default:
		fmt.Println("not equal")
	}

	// Output:
	// 10
}

func Example6() {
	switch {
	}

	// Output:
}
