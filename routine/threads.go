// In the current implementation of the runtime (Jan 2012) Go does not parallelize code by default,
// only a single core or processor is dedicated to a Go-program, regardless of how many goroutines
// are started in it; so these goroutines are running concurrent, they are not running in parallel: only one goroutine is running at a time.
// 现状已经不是上面的情况了
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 默认返回了8 == runtime.NumCPU()
	// GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting.
	fmt.Println(runtime.GOMAXPROCS(-1))
	fmt.Println(runtime.NumCPU())
}
