// So generally, using + operator to concatenate strings
// is convenient and efficient if the number of the concatenated
// strings is known at compile time.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func ExampleSprintf() {
	a := "hello"
	b := "world"
	fmt.Println(fmt.Sprintf("%s%s", a, b))
	// Output:
	// helloworld
}

func ExampleJoin() {
	a := "hello"
	b := "world"
	fmt.Println(strings.Join([]string{a, b}, ""))
	// Output:
	// helloworld
}

func ExampleBuffer() {
	a := "hello"
	b := "world"
	var v bytes.Buffer
	v.WriteString(a)
	v.WriteString(b)
	fmt.Println(v.String()) // 一次深度拷贝
	// Output:
	// helloworld
}

func ExampleBuilfer() {
	// Comparing with bytes.Buffer way, this way avoids
	// making an unnecessary duplicated copy of underlying
	// bytes for the resultant string.
	a := "hello"
	b := "world"
	var v strings.Builder
	v.WriteString(a)
	v.WriteString(b)
	fmt.Println(v.String())
	// Output:
	// helloworld
}
