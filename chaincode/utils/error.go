package utils

import "errors"

var (
	ErrNotExist          = errors.New("not existed")
	ErrAlreadyExist      = errors.New("already exist")
	ErrNotValid          = errors.New("not valid")
	ErrNotMatched        = errors.New("not matched")
	ErrCannotBeEmpty     = errors.New("can not be empty")
	ErrNotTransferable   = errors.New("not transferable")
	ErrAlreadyRegistered = errors.New("already registered")
)
