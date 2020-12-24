package bank

type bank struct {
	money    int
	deposits chan int
	withdraw chan int
	balances chan int
}

func NewBank(money int) *bank {
	b := bank{
		deposits: make(chan int),
		withdraw: make(chan int),
		balances: make(chan int),
	}
	b.money = money

	return &b
}

// 存钱
func (b *bank) Deposit(money int) { b.deposits <- money }

// 取钱
func (b *bank) Withdraw(money int) { b.withdraw <- money }

// 查看账户余额
func (b *bank) Balance() int { return <-b.balances }

func (b *bank) Run() {
	for {
		select {
		case money := <-b.deposits:
			b.money += money
		case money := <-b.withdraw:
			b.money -= money
		case b.balances <- b.money:
		}
	}
}
