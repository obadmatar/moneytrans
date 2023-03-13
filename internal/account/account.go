package account

import "github.com/shopspring/decimal"

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
