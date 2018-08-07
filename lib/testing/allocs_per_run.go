package main

import "testing"
import "fmt"

func test() {
	fmt.Println("testing")
	fmt.Println("testing")
	/* not a alloc */
	d := new(int)
	*d = 10
	fmt.Println(d)
}

func main() {
	fmt.Println(testing.AllocsPerRun(3, test))
	fmt.Println(testing.Coverage())
}
