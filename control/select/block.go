package main

import (
	"errors"
	"fmt"
)

type I interface{ error }

type J interface{}

func main() {

	err := errors.New("")

	var i I = err
	var j J = err

	fmt.Println(i == j)
}
