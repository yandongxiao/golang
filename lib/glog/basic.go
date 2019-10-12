// Package glog implements logging analogous to the Google-internal C++ INFO/ERROR/V setup.
// It provides functions Info, Warning, Error, Exit, Fatal, plus formatting variants such as Infof.
// NOTE: 没有Debug日志
package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// By default, all log statements write to files in a temporary directory.
	// This package provides several flags that modify this behavior.
	// As a result, flag.Parse must be called before any logging is done.
	flag.Parse()
	glog.Info("Prepare to repel boarders")
	// Log output is buffered and written periodically using Flush.
	// Programs should call Flush before exiting to guarantee all log output is written.
	defer glog.Flush()
}
