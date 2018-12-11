package main

import "fmt"

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(money int) { deposits <- money } // 只有一行语句的函数
func Balance() int      { return <-balances }

func broker() {
	money := 0 // 原本共享的变量，变成协程独有
	for {
		select {
		case inc := <-deposits:
			money += inc
		case balances <- money: // NOTICE: get money的方法. case语句为空
		}
	}
}

func main() {
	go broker()
	Deposit(100)
	fmt.Println(Balance())
}
