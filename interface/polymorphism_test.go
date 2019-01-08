// 一个struct类型的方法，最好是全是Pointer或全是对象本身.
package main

import "fmt"
import "math"

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
// 试着一个pointer 一个value receiver type的形式，golang会报错：
// rect does not implement geometry (area method has pointer receiver)
func (r *rect) area() float64 {
	return r.width * r.height
}
func (r *rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func ExamplePolymorphism() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(&r)
	measure(c)

	// Output:
	// &{3 4}
	// 12
	// 14
	// {5}
	// 78.53981633974483
	// 31.41592653589793
}
