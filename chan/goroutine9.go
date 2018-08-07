package main

import "time"

func send(ch chan int) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		_ = ch
		println("send data")
	}
}

func recv(ch chan int) {
	for {
		_, ok := <-ch
		if !ok {
			break
		}
		println("ok")
	}
}

func main() {
	ch := make(chan int)
	defer close(ch)
	go send(ch)
	go recv(ch)

	time.Sleep(100 * time.Second)
}
