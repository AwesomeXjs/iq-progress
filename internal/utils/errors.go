package utils

import "github.com/pkg/errors"

var (
	// ErrUserNotFound is returned when the specified user does not exist.
	ErrUserNotFound = errors.New("user not found")
	// ErrNotEnoughBalance is returned when the user does not have sufficient balance.
	ErrNotEnoughBalance = errors.New("not enough balance")
	// ErrSenderNotFound is returned when the sender does not exist in the transaction.
	ErrSenderNotFound = errors.New("failed executing code inside transaction: no rows in result set")
)
