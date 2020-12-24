package bank_test

import (
	"testing"

	"github.com/magiconair/properties/assert"

	bank "github.com/yandongxiao/go/chan/cases/bank/chan-bank"
)

func TestChan(t *testing.T) {
	b := bank.NewBank(0)
	go b.Run()

	b.Deposit(100)
	b.Deposit(100)
	b.Withdraw(10)

	assert.Equal(t, b.Balance(), 190)
}
