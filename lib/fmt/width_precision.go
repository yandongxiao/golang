// Width is specified by an optional decimal number immediately preceding the verb.
// If absent, the width is whatever is necessary to represent the value.
//
// Precision is specified after the (optional) width by a period followed by a decimal number.
// If no period is present, a default precision is used.
// A period with no following number specifies a precision of zero.

package main

import "fmt"

func main() {
	v := 12.3456

	// NOTE: For most values, width is the minimum number of runes to output, padding
	// the formatted form with spaces if necessary.
	fmt.Println("Width")
	fmt.Printf("%v\n", v)
	fmt.Printf("%9v\n", v)
	fmt.Printf("%3v\n", v)

	fmt.Println("Precision")
	fmt.Printf("% 9f\n", v) // 默认Precision=6
	fmt.Printf("%9.2f\n", v)
	fmt.Printf("%9.f\n", v) // A period with no following number specifies a precision of zero.
	fmt.Printf("%.2f\n", v) // 12.35
	fmt.Printf("%.3g\n", v) // 12, 注意与%.2f的区别

	fmt.Println("*")
	// Either or both of the flags may be replaced with the character '*'
	// causing their values to be obtained from the next operand, which must be of type int.
	fmt.Printf("%.*f\n", 3, v)

	// For strings, byte slices and byte arrays, however,
	// precision limits the length of the input to be formatted (not the size of the output), truncating if necessary.
	fmt.Println("the precision of string")
	fmt.Printf("%9.3s", "hello") // hel
}
