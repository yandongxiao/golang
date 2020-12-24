package bank

import "sync"

type bank struct {
	money int
	mutex sync.Mutex
}

func NewBank(money int) *bank {
	b := bank{money: money}
	return &b
}

// 存钱
func (b *bank) Deposit(money int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.money += money
}

// 取钱
func (b *bank) Withdraw(money int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.money -= money
}

// 查看账户余额
func (b *bank) Balance() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.money
}
