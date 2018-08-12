package main

import "os"
import "fmt"

func main() {
	fmt.Printf("%T\n", os.Args) /* []string */

	/* go run: /tmp/go-build972288702/command-line-arguments/_obj/exe/arguments */
	/* go build: */
	for i := range os.Args {
		fmt.Println(os.Args[i])
	}
}
