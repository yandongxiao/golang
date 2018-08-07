// struct 是 value type，所以使用new来创建新的对象
package main

import "fmt"

type Person struct {
	name string
}

func main() {
	// 四者等价
	p1 := &Person{"jack"}

	p2 := new(Person)
	p2.name = "jack"

	p3 := &Person{}
	p3.name = "jack"

	p4 := &Person{name: "jack"} // 推荐

	fmt.Printf("%T %v\n", p1, p1)
	fmt.Printf("%T %v\n", p2, p2)
	fmt.Printf("%T %v\n", p3, p3)
	fmt.Printf("%T %v\n", p4, p4)
}
