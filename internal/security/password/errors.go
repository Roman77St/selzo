package password

import "errors"

var (
	ErrPasswordTooShort = errors.New("password is too short")
	ErrPasswordTooWeak  = errors.New("password is too weak")
)