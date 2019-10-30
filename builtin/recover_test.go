package main

import "fmt"

func ExampleBasic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("hello")
	fmt.Println("world")
	//Output:
	//hello
}

func ExampleNil1() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println(recover())
	//Output:
	//<nil>
	//<nil>
}

func ExampleNil2() {
	defer func() {
		fmt.Println("hello")
	}()
	defer func() {
		fmt.Println(recover())
	}()
	panic(nil)
	//Output:
	// <nil>
	// hello
}
