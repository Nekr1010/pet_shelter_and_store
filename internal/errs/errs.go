package errs

import "errors"

var (
	ErrNoPermissionsToWithdraw = errors.New("no permissions to withdraw this account")
	ErrUserIDNotFound          = errors.New("user ID Not Found")
	ErrUserAlreadyExists       = errors.New(`user already exists`)
	ErrNotFound                = errors.New("not found")
	ErrNotEnoughBalance        = errors.New("not enough balance")
	ErrInvalidOperationType    = errors.New("invalid operation type")
)
