package main

import (
	"bytes"
	"fmt"
)

/* Due to compiler-optimizations and depending on the size of the strings using a
Buffer only starts to become more efficient when the number of iterations > 15.
*/
func main() {
	var buff bytes.Buffer
	buff.WriteString("helloworld\n")
	buff.WriteString("helloworld\n")
	buff.WriteString("helloworld\n")
	buff.WriteString("你好")
	fmt.Println(buff.String())
}
