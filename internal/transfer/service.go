package transfer

import (
	"sync"

	"github.com/obadmatar/moneytrans/internal/account"
	"github.com/shopspring/decimal"
)

type Service struct {
	mutex sync.Mutex
	repo  account.Repositoy
}

type ServiceConfig func(s *Service) error

func NewService(configs...ServiceConfig) (*Service, error) {
	s := &Service{}

	for _, cfg := range configs {
		err := cfg(s)

		if err != nil {
			return nil, err
		}
	}

	return s, nil
}

func WithInMemoryAccountRepository(filepath string) ServiceConfig {
	return func(s *Service) error {
		s.repo = account.NewInMemoryRepository(filepath)
		return nil
	}
}

// TODO: should be hanlded using transactions to ensure atomicity of updates in case of any failures.

// Transfer moves the specified amount from one account to another.
// it ensures that the 'fromAccount' has sufficient balance and updates the account balances accordingly
func (s *Service) Transfer(fromAccountID, toAccountID string, amount decimal.Decimal) error {
	// Acquire lock before accessing account balances
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check if amount is not negative
	if !amount.IsPositive() {
		return account.ErrNegativeAmount
	}

	// Get accounts
	fromAccount, err := s.repo.Get(fromAccountID)
	if err != nil {
		return err
	}
	toAccount, err := s.repo.Get(toAccountID)
	if err != nil {
		return err
	}

	// Check if from account has sufficient balance
	if fromAccount.Balance().LessThan(amount) {
		return account.ErrNegativeBalance
	}

	// Update account balances
	fromAccount.Debit(amount)
	toAccount.Credit(amount)

	// Save updates
	err = s.repo.Update(fromAccount)
	if err != nil {
		return err
	}
	err = s.repo.Update(toAccount)
	if err != nil {
		return err
	}

	return nil
}
