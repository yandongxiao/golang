package main

type Person struct{}

// NOTICE: 必须去掉p
func (p *Person) Add(a, b int) int {
	return a + b
}

var foo func(p *Person, a, b int) int = (*Person).Add

func main() {
	println(foo(nil, 1, 2))

	var p Person
	var bar func(x, y int) int = p.Add // NOTICE: 不一定传递p的指针
	println(bar(1, 2))
}
