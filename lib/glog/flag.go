package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	// By default, all log statements write to files in a temporary directory.
	// This package provides several flags that modify this behavior.
	// As a result, flag.Parse must be called before any logging is done.
	// -logtostderr=false：Logs are written to standard error instead of to files.
	// -alsologtostderr=false: Logs are written to standard error as well as to files.
	// -stderrthreshold=ERROR: Log events at or above this severity are logged to standard error.
	//                  go run basic.go -log_dir='/tmp' -stderrthreshold=INFO 此时日志也会输出的标准错误
	//					NOTE: 如果-logtostderr或-alsologtostderr被设置，那么一定会输出到标准错误。
	//					NOTE: 该标志是无法控制INFO级别的日志，不输出到日志文件的。除非使用重定向功能, 将标准错误内容进行重定向.
	// -log_dir="" : 注意带引号, Log files will be written to this directory instead of the default temporary directory.
	//               文件名称由glog决定
	// -log_backtrace_at=basic.go:17 该位置是glog的打印语句，则打印堆栈信息
	flag.Parse()
	glog.Info("Prepare to repel boarders")
	defer glog.Flush()
}
