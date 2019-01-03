package main

import "fmt"

var (
	token   = make(chan struct{}, 1)
	account int
)

func TDeposit(money int) {
	token <- struct{}{}
	account += money
	<-token
}

func TBalance() int {
	token <- struct{}{}
	val := account
	<-token
	return val
}

func ExampleToken() {
	TDeposit(100)
	fmt.Println(TBalance())
	//Output:
	//100
}
