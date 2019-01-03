package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func MDeposit(account int) {
	mutex.Lock()
	defer mutex.Unlock()
	balance += account
}

func MBalance() int {
	mutex.Lock()
	defer mutex.Unlock()
	return balance
}

func ExampleMutex() {
	MDeposit(100)
	fmt.Println(MBalance())
	//Output:
	//100
}
