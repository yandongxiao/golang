// only the gc-compilers have a true implementation of goroutines, mapping them onto OS threads as appropriate.
// With the gccgo compiler, an OS thread will be created for each goroutine.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Compiler)
}
