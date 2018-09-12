// Package log implements a **simple** logging package.
// It defines a type, Logger, with methods for formatting output.
// Every log message is output on a separate line: if the message being printed
// does not end in a newline, the logger will add one.
// NOTE: Each logging operation makes a single call to the Writer's Write method.
// NOTE: A Logger can be used simultaneously from multiple goroutines; it guarantees to serialize access to the Writer.
//       writer可以不是线程安全的
package main

import (
	"log"
	"os"
)

func main() {
	//standardLogger()
	newLogger()
}

func newLogger() {
	f, _ := os.Create("/tmp/output")
	defer f.Close()
	logger := log.New(f, "", 0)
	// The Fatal functions call os.Exit(1) after writing the log message
	logger.Fatalf("fatal errror=%v", "testing")
}

// It also has a predefined 'standard' Logger accessible through helper functions
// Print[f|ln], Fatal[f|ln], and Panic[f|ln].
// That logger writes to standard error and prints the date and time of each logged message.
func standardLogger() {
	// The Panic functions call panic after writing the log message.
	// panic()的参数值是格式化的字符串
	log.Panicf("fatal errror=%v", "testing")
}
