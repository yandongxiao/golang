package bank

import "sync"

var (
	mutex   sync.Mutex
	balance int
)

func Deposit(account int) {
	mutex.Lock()
	defer mutex.Unlock()
	balance += account
}

func Balance() int {
	mutex.Lock()
	defer mutex.Unlock()
	return balance
}
