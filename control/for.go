// for有四种形式, 并且是遍历的唯一方式
package main

import "fmt"

func main() {
	// for 是golang当中唯一的循环遍历的方式
	// 支持continue 和 break
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", i)
	}

	for {
		println("while true")
		break
	}

	i := 1
	for i < 10 {
		fmt.Printf("%d, ", i)
		i++
	}

	for i, v := range []int{1, 2, 3} {
		println(i, v)
	}

	// for range 简化形式
	for range []int{1, 2, 3} {
		println("--")
	}

	// NOTE: range 的参数值可以是nil! 但不支持 for range nil
	var strs []string
	for i := range strs {
		println(strs[i])
	}

	// 自动解引用
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array)
	fmt.Printf("The sum of the array is: %f", x)
}

// can also with dereferencing *a to get back to the array
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
