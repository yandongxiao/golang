package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%c\n", os.PathSeparator)
	fmt.Printf("%c\n", os.PathListSeparator)
}
