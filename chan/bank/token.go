package bank

var (
	token   = make(chan struct{}, 1)
	account int
)

func Deposit(money int) {
	token <- struct{}{}
	account += money
	<-token
}

func Balance() int {
	token <- struct{}{}
	val := account
	<-token
	return val
}
