package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Package signal will not block sending to c: the caller must ensure that
	// c has sufficient buffer space to keep up with the expected signal rate.
	// For a channel used for notification of just one signal value, a buffer of size 1 is sufficient.
	// 如果chan处理时间太长，chan buffer塞满了信号，新来的信号会丢失哟
	ch1 := make(chan os.Signal)
	ch2 := make(chan os.Signal)
	signal.Notify(ch1, syscall.SIGINT) // ctl-c
	signal.Notify(ch1, syscall.SIGTSTP)
	signal.Notify(ch2, syscall.SIGINT)
	signal.Notify(ch2, syscall.SIGTSTP) // ctl-z

	go func(ch <-chan os.Signal) {
		for sig := range ch {
			switch sig {
			case syscall.SIGINT:
				time.Sleep(time.Second)
				fmt.Println("syscall.SIGINT")
				os.Exit(1)
			case syscall.SIGTSTP:
				fmt.Println("syscall.SIGTSTP")
				time.Sleep(time.Second)
				fmt.Println("restart")
			}
		}
	}(ch1)

	go func(ch <-chan os.Signal) {
		for sig := range ch {
			switch sig {
			case syscall.SIGINT:
				fmt.Println("..")
			case syscall.SIGTSTP:
				fmt.Println("..")
			}
		}
	}(ch2)

	time.Sleep(time.Second * 5)
}
