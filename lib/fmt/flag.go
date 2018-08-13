// Flags are ignored by verbs that do not expect them.
package main

import "fmt"

func main() {
	// always print a sign for numeric values
	fmt.Printf("%+6d\n", 123)

	// pad with spaces on the right rather than the left
	fmt.Printf("%-6d\n", 123)

	// 0 for octal (%#o), 0x for hex (%#x); 0X for hex (%#X)
	fmt.Printf("%#x\n", 123)

	fmt.Printf(" %d% d\n", 1, 2)       // 打印 1 2
	fmt.Printf("% #x\n", "helloworld") // 0x68 0x65 0x6c 0x6c 0x6f 0x77 0x6f 0x72 0x6c 0x64

	// pad with leading zeros rather than spaces
	fmt.Printf("%010d\n", 123)
}
