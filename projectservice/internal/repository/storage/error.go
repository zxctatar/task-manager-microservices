package storage

import "errors"

var (
	ErrAlreadyExists = errors.New("entry alredy exists")
	ErrNotFound      = errors.New("entry not found")
)
