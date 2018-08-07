package main

import "fmt"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := fmt.Errorf("unkonw error")
	errcheck(err)
}
