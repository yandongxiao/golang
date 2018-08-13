// package fmt的函数主要包括两大类：print and scan
// scan包括:  Scan, Scanf, Scanln,		Fscan, Fscanf, Fscanln,		Sscan, Sscanf, Sscanln
// print包括：Print, Printf, Println,	Fprint, Fprintf, Fprintln,	Sprint, Sprintf, Sprintln
// fmt.Errorf用于格式化一个error
package main

import "fmt"

func main() {
	err := fmt.Errorf("%s%s", "hello", "world")
	fmt.Println(err)

	var a, b string
	fmt.Sscanf("hello world", "%s%s", &a, &b)
	fmt.Println(a, b)
}
