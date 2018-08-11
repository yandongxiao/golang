package main

import "fmt"
import "unicode/utf8"

func test1() {
	var str = "国"
	vals := []byte(str)
	vals[2]++
	fmt.Printf("%s\n", vals)
}

func main() {
	var str = "国"
	fmt.Println(str)
	fmt.Println("len is", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d\n", str[i])
	}

	//str[2] += 1   /* you can not change string */
	for _, val := range str {
		fmt.Printf("%c\n", val)
	}
	println(str)

	println(utf8.RuneCountInString(str))
	test1()
}
