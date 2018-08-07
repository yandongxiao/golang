// 理解什么是一个receiver ?
// recevier是一个type，所以Shape或Rectangle可以作为一个receiver
// receiver不能是一个指针，所以type P *Rectangle定义的新类型P，不能作为一个receiver
// (r *Rectangle) Area 和 (r *Rectangle) Area 不能同时定义, 否则会报告重定义

package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Shape struct{}

func (sh Shape) Area() float32 {
	return -1 // the shape is indetermined, so we return something impossible
}

type Rectangle struct {
	length, width float32
	Shape
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

// method redeclared: Rectangle.Area
/*
func (r Rectangle) Area() float32 {
	return 1.1
}
*/

func main() {
	s := Shape{}
	r := Rectangle{5, 3, s} // Area() of Rectangle needs a value
	shapes := []Shaper{r, s}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}

/* Output:
Looping through shapes for area ...
Shape details:  {5 3}
Area of this shape is:  15
Shape details:  {}
Area of this shape is:  -1
*/
