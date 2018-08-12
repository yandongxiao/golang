package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ch1 := make(chan os.Signal)
	signal.Notify(ch1, syscall.SIGINT)
	signal.Stop(ch1)
	signal.Notify(ch1, syscall.SIGINT) // 执行有效

	// Stop causes package signal to stop relaying incoming signals to c.
	// 但是发送端并没有关闭(因为下面的语句被阻塞了)
	fmt.Println(<-ch1)

	time.Sleep(time.Second)
}
