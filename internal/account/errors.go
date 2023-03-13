package account

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrNegativeBalance = errors.New("nigative balance")
	ErrNegativeAmount = errors.New("nigative amount")
)