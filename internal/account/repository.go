package account

import (
	"encoding/json"
	"os"
	"sync"
)

type Repositoy interface {
	List() ([]*Account, error)

	Get(id string) (*Account, error)

	Update(account *Account) error
}

// inMemoryRepository is an in-memory implementation of the Repository interface
type inMemoryRepository struct {
	mutex    sync.Mutex
	accounts map[string]Account
}

func NewInMemoryRepository(filePath string) Repositoy {
	// read accounts from JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// unmarshal JSON
	var accounts []*Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		panic(err)
	}

	// init accounts map
	accountsMap := make(map[string]Account)
	for _, acc := range accounts {
		accountsMap[acc.id] = *acc
	}

	return &inMemoryRepository{accounts: accountsMap}
}

// List returns a list of all accounts
func (r *inMemoryRepository) List() ([]*Account, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	accounts := make([]*Account, 0, len(r.accounts))

	for _, acc := range r.accounts {
		accounts = append(accounts, &acc)
	}

	return accounts, nil
}

// Get returns the account with the specified ID
func (r *inMemoryRepository) Get(id string) (*Account, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if acc, ok := r.accounts[id]; ok {
		return &acc, nil
	}

	return nil, ErrAccountNotFound
}

// Update updates the specified account
func (r *inMemoryRepository) Update(account *Account) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.accounts[account.id]; ok {
		r.accounts[account.id] = *account
		return nil
	}

	return ErrAccountNotFound
}
