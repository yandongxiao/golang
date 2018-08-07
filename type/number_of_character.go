package main

import "unicode/utf8"

func main() {
	name := "你好"
	/* cannot convert name (type string) to type []int */
	//println(len([]int(name)))
	println(utf8.RuneCountInString(name))
}
