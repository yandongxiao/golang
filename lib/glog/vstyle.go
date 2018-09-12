// go run vstyle.go -logtostderr=true -stderrthreshold=INFO -v=10
// go run vstyle.go -logtostderr=true -vmodule="v*=10"
// v-style的日志工作流貌似是这样的.
//     1. 某条日志是否被打印不再取决于一开始设置的日志级别, -stderrthreshold.
//	   2. 通过比较两个数字的大小来决定是否打印日志。-v=10(相当于指定了-stderrthreshold.)
//        代码中，glog.V(10).Infoln的日志将会被打印, [0, 10]都是可以的.
//     3. V-style最大的特点是，glog.V(level) level可以以变量为输入。可以不停服的情况下，更新level值
package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	level()
}

func basic() {
	// -v=0: Enable V-leveled logging at the specified level.
	// -vmodule=gopher*=3: sets the V level to 3 in all Go files whose names begin "gopher"
	// The syntax of the argument is a comma-separated list of pattern=N, where pattern is
	// a literal file name (minus the ".go" suffix) or "glob" pattern and N is a V level.
	flag.Parse()

	glog.V(9).Infoln("just info message") // 也可以直接传递0常量
	glog.Infoln("helloworld")
}

func level() {
	flag.Parse()
	// the -v flag is of type Level and should be modified only through the flag.Value interface.
	// Level specifies a level of verbosity for V logs. *Level implements flag.Value(这样)
	// 通过level的Set接口，即可修改-v指定的Log Level值
	level := new(glog.Level)
	level.Set("9") //level本身的值不再重要

	glog.V(9).Infoln("just info message")
}
