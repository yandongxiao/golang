// Bits or'ed together to control what's printed.
// There is no control over the order they appear or the format they present
package main

import (
	"log"
	"os"
)

func main() {
	// prefix为空，flag为0，则logger的输出内容与fmt输出内容相同
	// Ldate: 2009/01/23
	// Ltime: 01:23:23
	// Lmicroseconds: 123123
	// Llongfile: /a/b/c/d.go:23
	// Lshortfile: d.go:23
	logger := log.New(os.Stdout, "--", 0)
	logger.Printf("hello")
	logger.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	logger.Printf("world")
}
