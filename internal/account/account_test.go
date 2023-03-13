package account

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestCredit(t *testing.T) {
	t.Run("Test Account Credit", func(t *testing.T) {
		acc := &Account{
			balance: decimal.NewFromInt(100),
		}
		acc.Credit(decimal.NewFromInt(100))
		expected := decimal.NewFromInt(200)
		if !acc.Balance().Equal(expected) {
			t.Errorf("Balance: expected %s, got %s", expected, acc.Balance())
		}
	})

	t.Run("Test Account Credit With No Loss of Precision", func(t *testing.T) {
		acc := &Account{
			balance: decimal.NewFromInt(100),
		}
		acc.Credit(decimal.NewFromFloat(99.9))
		expected := decimal.NewFromFloat(199.9)
		if !acc.Balance().Equal(expected) {
			t.Errorf("Balance: expected %s, got %s", expected, acc.Balance())
		}
	})
}

func TestDebit(t *testing.T) {
	t.Run("Test Account Debit", func(t *testing.T) {
		acc := &Account{
			balance: decimal.NewFromInt(100),
		}
		acc.Debit(decimal.NewFromInt(1))
		expected := decimal.NewFromInt(99)
		if !acc.Balance().Equal(expected) {
			t.Errorf("Balance: expected %s, got %s", expected, acc.Balance())
		}
	})

	t.Run("Test Account Debit With No Loss of Precision", func(t *testing.T) {
		acc := &Account{
			balance: decimal.NewFromInt(100),
		}
		acc.Debit(decimal.NewFromFloat(0.1))
		expected := decimal.NewFromFloat(99.9)
		if !acc.Balance().Equal(expected) {
			t.Errorf("Balance: expected %s, got %s", expected, acc.Balance())
		}
	})
}
