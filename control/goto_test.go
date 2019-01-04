package main

import "fmt"

func ExampleGoTo() {
	gotoCount := 0
GotoLabel:
	gotoCount++
	if gotoCount < 10 {
		goto GotoLabel //如果小于10的话就跳转到GotoLabel
	}
	fmt.Println(gotoCount)
	// Output:
	// 10
}

func ExampleBad1() {
	v := 3
	goto L // BAD
	// ./goto_test.go:15:7: goto L jumps over declaration of v at ./goto_test.go:16:4
	// 导致编译通不过, 所以必须上移该语句至goto之上
	// v := 3
	v++
L:
	fmt.Println(v)

	//Output:
	//3
}

// 原则：A "goto" statement outside a block cannot jump to a label inside that block.
// ./goto_test.go:31:8: goto L1 jumps into block starting at ./goto_test.go:33:12
// 导致编译通不过，所以必须注释掉goto语句
// label L1 defined and not used. NOTE：也不能随意添加label
// 导致编译通不过，所以必须去掉注释L1标签
func ExampleBad2() {
	n := 10
	if n%2 == 1 {
		//goto L1
	}

	if n > 0 {
		n--
		// L1:
		n--
	}
}
