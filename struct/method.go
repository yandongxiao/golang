package main

import "fmt"

// method value: have been bound to a specific receiver value. e.g. w.Write
// func (p []byte) (n int, err error) {
//	return w.Write(p)
// }]
//
// method expressions: generate functions from methods of a given type, e.g. (*bufio.Writer).Write
// func (w *bufio.Writer, p []byte) (n int, err error) {
//		return w.Write(p)
// }])

type Person struct {
	name string
}

func (p Person) Name() string {
	return p.name
}

func main() {
	p := Person{
		name: "jack",
	}

	//  The expression p is evaluated and saved during the evaluation of the method value;
	// the saved **copy** is then used as the receiver in any calls, which may be executed later.
	foo := p.Name
	p.name = "alice"
	fmt.Println(foo()) // 解释为啥不是alice
}
