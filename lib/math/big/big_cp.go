package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	v1 := big.NewInt(math.MaxInt64)
	v2 := big.NewInt(1)
	fmt.Println(v1)
	fmt.Println(v2)
	v2.Add(v1, v2)
	fmt.Println(v1)
	fmt.Println(v2)

	vv1 := big.NewRat(12, 100)
	vv2 := big.NewRat(88, 100)
	vv1.Add(vv1, vv2)
	fmt.Println(vv1)
}
