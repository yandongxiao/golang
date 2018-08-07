package main

import (
	"fmt"
	"time"
)

type Person struct {
	age    int
	handle chan bool
}

func (p *Person) work() {
	for i := 0; i < 100; i++ {
		<-p.handle // 接收句柄
		p.age++
		p.handle <- true // 释放句柄
	}
}

func main() {
	p := Person{0, make(chan bool)}
	for i := 0; i < 100; i++ {
		go p.work()
	}
	p.handle <- true
	time.Sleep(time.Second)
	fmt.Println(p.age)
}
