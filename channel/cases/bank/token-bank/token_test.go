package bank_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	bank "github.com/yandongxiao/go/chan/cases/bank/token-bank"
)

func TestChan(t *testing.T) {
	b := bank.NewBank(0)

	b.Deposit(100)
	b.Deposit(100)
	b.Withdraw(10)

	assert.Equal(t, b.Balance(), 190)
}
