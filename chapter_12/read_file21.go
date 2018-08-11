package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input, err := os.Open("products2.txt")
	if err != nil {
		panic("can not open file")
	}

	var v1, v2, v3 []string
	for {
		var a, b, c string
		_, err := fmt.Fscanln(input, &a, &b, &c)
		if err == io.EOF {
			break
		}

		v1 = append(v1, a)
		v2 = append(v2, b)
		v3 = append(v3, c)
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
}
