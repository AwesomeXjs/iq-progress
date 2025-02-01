package utils

import "github.com/pkg/errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrNotEnoughBalance = errors.New("not enough balance")
	ErrSenderNotFound   = errors.New("failed executing code inside transaction: no rows in result set")
)
