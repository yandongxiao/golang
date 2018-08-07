package main

import (
	"fmt"
	"time"
)

func main() {
	// 一旦创建timer以后，timer就开始工作，并在2小时以后触发事件，并向timer.C中写时间
	timer := time.NewTimer(time.Second)
	<-timer.C
	fmt.Println("timer expired")

	timer = time.NewTimer(time.Second)
	go func() {
		<-timer.C
		fmt.Println("timer expired")
	}()
	if stoped := timer.Stop(); stoped {
		fmt.Println("timer is stoped before expired")
	}
}
