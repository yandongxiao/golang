package main

import "strconv"

func main() {
	println(strconv.IntSize)
	ival, err := strconv.Atoi("100")
	println(ival, err)
	fval, err := strconv.ParseFloat("3.2", 64)
	println(fval, err)

	println(strconv.FormatFloat(1.2, 'g', 8, 64))
}
