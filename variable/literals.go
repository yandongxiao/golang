// a composite literal of array, slice, or map type can elide the type specification for the elements' initializers
// if they are not pointer type.
package main

import "fmt"

type Date struct {
	month string
	day   int
}

func main() {
	// Struct values, fully qualified; always legal.
	holiday1 := []Date{
		Date{"Feb", 14},
		Date{"Nov", 11},
		Date{"Dec", 25},
	}
	// Struct values, type name elided; always legal.
	holiday2 := []Date{
		{"Feb", 14},
		{"Nov", 11},
		{"Dec", 25},
	}
	// Pointers, fully qualified, always legal.
	holiday3 := []*Date{
		&Date{"Feb", 14},
		&Date{"Nov", 11},
		&Date{"Dec", 25},
	}

	// Pointers, type name elided; legal in Go 1.
	// NOTE: 未来可能不支持
	holiday4 := []*Date{
		{"Feb", 14},
		{"Nov", 11},
		{"Dec", 25},
	}

	fmt.Println(holiday1, holiday2, holiday3, holiday4)
}
