package main

import (
	"fmt"
	"strings"
)

func main() {

	data := []string{
		"hello", "world",
	}
	fmt.Println(strings.Join(data, "/"))

	builder := strings.Builder{}
	builder.WriteString("hello")
	builder.WriteString("world")
	fmt.Println(builder.String())
}
