package transfer

import (
	"testing"

	"github.com/obadmatar/moneytrans/internal/account"
	"github.com/shopspring/decimal"
)

func TestTransfer(t *testing.T) {
	r := account.NewInMemoryRepository("../account/data/accounts_test.json")
	s := Service{repo: r}

	fromAccountId := "3d253e29-8785-464f-8fa0-9e4b57699db9" //  balance = 87.11
	toAccountId := "17f904c1-806f-4252-9103-74e7a5d3e340"   //  balance = 946.15

	t.Run("Test Transferring An Amount Greater Than The FromAccount Balance", func(t *testing.T) {
		err := s.Transfer(fromAccountId, toAccountId, decimal.NewFromFloat(150.0))
		if err != account.ErrNegativeBalance {
			t.Errorf("Expected ErrNegativeBalance error, but got %v", err)
		}
	})

	t.Run("Test Transferring An Amount Less Than The FromAccount Balance", func(t *testing.T) {
		err := s.Transfer(fromAccountId, toAccountId, decimal.NewFromFloat(7.11))
		if err != nil {
			t.Errorf("Expected nil error, but got %v", err)
		}
		updatedFromAcc, _ := r.Get(fromAccountId)
		updatedToAcc, _ := r.Get(toAccountId)
		expectedFromBalance := decimal.NewFromFloat(80.0)
		expectedToBalance := decimal.NewFromFloat(953.26)
		if updatedFromAcc.Balance().Cmp(expectedFromBalance) != 0 {
			t.Errorf("Expected %v balance for from account, but got %v", expectedFromBalance, updatedFromAcc.Balance())
		}
		if updatedToAcc.Balance().Cmp(expectedToBalance) != 0 {
			t.Errorf("Expected %v balance for to account, but got %v", expectedToBalance, updatedToAcc.Balance())
		}
	})

	  t.Run("Test Transferring From A Non Existent Account", func(t *testing.T) {
        err := s.Transfer("non-exsistent", toAccountId, decimal.NewFromFloat(50.0))
        if err != account.ErrAccountNotFound {
            t.Errorf("Expected ErrAccountNotFound error, but got %v", err)
        }
    })


	  t.Run("Test Transferring To A Non Existent Account", func(t *testing.T) {
        err := s.Transfer(fromAccountId, "non-exsistent", decimal.NewFromFloat(50.0))
        if err != account.ErrAccountNotFound {
            t.Errorf("Expected ErrAccountNotFound error, but got %v", err)
        }
    })

	t.Run("Test Transferring Negative Amount", func(t *testing.T) {
        err := s.Transfer(fromAccountId, fromAccountId, decimal.Zero.Neg())
        if err != account.ErrNegativeAmount {
			t.Errorf("Expected ErrNegativeAmount error, but got %v", err)
		}
    })
}
