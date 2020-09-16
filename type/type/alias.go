package main

import "fmt"

type INT = int

// cannot define new methods on non-local type int
//func (i INT) String() string {
//	return fmt.Sprintf("%d", i)
//}

func main() {
	a := INT(10)
	fmt.Println(a)
}
