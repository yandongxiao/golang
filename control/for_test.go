// for有四种形式, 并且是遍历的唯一方式
package main

import "fmt"

func ExampleFormat1() {
	// for 是golang当中唯一的循环遍历的方式
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	// Output:
	// 0123456789
}

func ExampleFormat2() {
	i := 1
	for i < 10 {
		fmt.Printf("%d", i)
		i++
	}
	// Output:
	// 123456789
}

func ExampleFormat3() {
	for {
		fmt.Println("helloworld")
	}
	// Output
	// helloworld
}

func ExampleForRange1() {
	// 支持map, slice, array, pointer
	for i, v := range []int{1, 2, 3} {
		fmt.Printf("%d %d, ", i, v)
	}
	// Output
	// 0 1, 1 2, 2 3
}

func ExampleForRange2() {
	// for range 简化形式
	for range []int{1, 2, 3} {
		print("--")
	}
	// Output
	// ------
}

func ExampleNil() {
	// NOTE: range 的参数值可以是nil! 但不支持 for range nil
	var strs []string
	for i := range strs {
		println(strs[i])
	}
	// Output
	//
}

func ExampleChan() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 1
	close(ch)
	for v := range ch {
		fmt.Printf("%d", v)
	}
	//Output:
	// 121
}

func ExamplePointer() {
	// 自动解引用
	// can also with dereferencing *a to get back to the array
	// NOTE: 即使p==nil, len(p) 仍然等于3，所以，不要使用i := 0; i < len(p); i++的形式，遍历指向数组的指针
	p := &[3]float64{7.0, 8.5, 9.1}
	sum := float64(0)
	for _, v := range p {
		sum += v
	}
	fmt.Printf("%f", sum)
	// Output
	// 24.6
}
