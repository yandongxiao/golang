package main

import "bytes"
import "fmt"

func main() {
	var data bytes.Buffer

	data.WriteString("aaa")
	data.WriteString("bbb")
	data.WriteString("ccc")

	fmt.Println(data.String())
}
