package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// 1. 设置子进程的属性
	procAttr := &os.ProcAttr{
		Dir: "/tmp",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	// 2. 启动子进程
	// 第一个参数必须使用全路径
	argv := []string{"/bin/ls", "/tmp"}
	proc, err := os.StartProcess(argv[0], argv, procAttr)
	checkError(err)

	// 3. 等待子进程调用完毕
	// Wait releases any resources associated with the Process. On most operating systems,
	// the Process must be a child of the current process or an error will be returned.
	stat, err := proc.Wait()
	checkError(err)
	fmt.Println(stat.Exited())
	fmt.Println(stat.Pid())
	fmt.Println(stat.String())
	fmt.Println(stat.Success())
	fmt.Println(stat.SysUsage())
	fmt.Println(stat.SystemTime())
	fmt.Println(stat.UserTime())
}
