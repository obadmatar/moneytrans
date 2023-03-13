package account

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Account struct {
	id      string
	name    string
	balance decimal.Decimal
}

func (acc *Account) Id() string {
	return acc.id
}

func (acc *Account) Name() string {
	return acc.name
}

func (acc *Account) Balance() decimal.Decimal {
	return acc.balance
}

// Credit adds the given amount to the account's balance.
func (acc *Account) Credit(amount decimal.Decimal) {
	acc.balance = acc.balance.Add(amount)
}

// Debit subtracts the given amount from the account's balance.
func (acc *Account) Debit(amount decimal.Decimal) {
	acc.balance = acc.balance.Sub(amount)
}

// UnmarshalJSON implements json.Unmarshaler interface
func (acc *Account) UnmarshalJSON(data []byte) error {
	type accountJSON struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Balance string `json:"balance"`
	}
	var aux *accountJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	balance, err := decimal.NewFromString(aux.Balance)
	if err != nil {
		return err
	}
	acc.id = aux.Id
	acc.name = aux.Name
	acc.balance = balance
	return nil
}
