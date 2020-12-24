package bank

type bank struct {
	money int
	token chan struct{}
}

func NewBank(money int) *bank {
	b := bank{money: money, token: make(chan struct{}, 1)}
	return &b
}

// 存钱
func (b *bank) Deposit(money int) {
	b.token <- struct{}{}
	defer func() {
		<-b.token
	}()
	b.money += money
}

// 取钱
func (b *bank) Withdraw(money int) {
	b.token <- struct{}{}
	defer func() {
		<-b.token
	}()
	b.money -= money
}

// 查看账户余额
func (b *bank) Balance() int {
	b.token <- struct{}{}
	defer func() {
		<-b.token
	}()
	return b.money
}
