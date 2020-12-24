// kill -SIGUSR1 56435
package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func ExampleA() {
	setupSigusr1Trap()
	go aa()
	m11()
}

func m11() {
	m22()
}

func m22() {
	m33()
}

func m33() {
	time.Sleep(time.Hour)
}

func aa() {
	time.Sleep(time.Hour)
}

func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			DumpStacks()
		}
	}()
}

func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, false)] // true 表示获取所有协程的堆栈信息
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
