package main

import "fmt"

func main() {
	from := []int{1, 2, 3}
	to := make([]int, 10)

	n := copy(to, from)
	fmt.Println(to)                       // output: [1 2 3 0 0 0 0 0 0 0]
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3) // output: [1 2 3 4 5 6]
}
