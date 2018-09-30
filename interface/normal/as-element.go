// 将interface type视为一种普通的类型
package main

import "fmt"

// interface value as an element of slice
func method1() {
	ds := []interface{}{}
	ds = append(ds, 10)
	ds = append(ds, "jack")
	ds = append(ds, 1.0)

	for _, elmt := range ds {
		switch t := elmt.(type) {
		case int:
			fmt.Println("int: ", t)
		case float64:
			fmt.Println("float64: ", t)
		case string:
			fmt.Println("string: ", t)
		}
	}
}

// interface value as an element of chan
func method2() {
	ds := make(chan interface{})
	go func() {
		ds <- "jack"
		ds <- 10
		close(ds)
	}()

	for elmt := range ds {
		fmt.Println(elmt)
	}
}

func main() {
	method1()
	method2()
}
