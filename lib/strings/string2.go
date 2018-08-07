package main

import "strings"

func main() {
	var str = "nihao"
	b := strings.Contains(str, "hao")
	if b {
		idx := strings.Index(str, "hao")
		idx = strings.LastIndex(str, "hao")
		println(idx)
	}

	strings.Replace(str, "hao", "ver", 1)
	println(str)

}
