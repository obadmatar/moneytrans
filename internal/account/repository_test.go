package account

import (
	"errors"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestInMemoryRepository(t *testing.T) {
	repo := NewInMemoryRepository("accounts_test.json")

	// Test List function
	t.Run("Test List function", func(t *testing.T) {
		accounts, err := repo.List()
		if err != nil {
			t.Errorf("Error listing accounts: %v", err)
		}
		if len(accounts) != 3 {
			t.Errorf("Expected 3 accounts, but got %d", len(accounts))
		}
	})

	// Test Get function with existing account
	t.Run("Test Get function with existing account", func(t *testing.T) {
		id := "3d253e29-8785-464f-8fa0-9e4b57699db9"
		expectedAccount := &Account{
			id:      id,
			name:    "Trupe",
			balance: decimal.NewFromFloat(87.11),
		}
		account, err := repo.Get(id)
		if err != nil {
			t.Errorf("Error getting account: %v", err)
		}
		if !reflect.DeepEqual(account, expectedAccount) {
			t.Errorf("Expected account to be %v, but got %v", expectedAccount, account)
		}
	})

	// Test Get function with non-existing account
	t.Run("Test Get function with non-existing account", func(t *testing.T) {
		account, err := repo.Get("non-existing-id")
		if !errors.Is(err, ErrAccountNotFound) {
			t.Errorf("Expected error %v, but got %v", ErrAccountNotFound, err)
		}
		if account != nil {
			t.Errorf("Expected account to be nil, but got %v", account)
		}
	})

	// Test Update function with existing account
	t.Run("Test Update function with existing account", func(t *testing.T) {
		id := "3d253e29-8785-464f-8fa0-9e4b57699db9"
		updatedAccount := &Account{
			id:      id,
			name:    "New Account",
			balance: decimal.Zero,
		}
		err := repo.Update(updatedAccount)
		if err != nil {
			t.Errorf("Error updating account: %v", err)
		}
		account, err := repo.Get(id)
		if err != nil {
			t.Errorf("Error getting account after update: %v", err)
		}
		if !reflect.DeepEqual(account, updatedAccount) {
			t.Errorf("Expected account to be %v, but got %v", updatedAccount, account)
		}
	})

	// Test Update function with non-existing account
	t.Run("Test Update function with non-existing account", func(t *testing.T) {
		updatedAccount := &Account{
			id:      "non-existing-id",
			name:    "Invalid Account",
			balance: decimal.Zero,
		}
		err := repo.Update(updatedAccount)
		if !errors.Is(err, ErrAccountNotFound) {
			t.Errorf("Expected error %v, but got %v", ErrAccountNotFound, err)
		}
	})

}
