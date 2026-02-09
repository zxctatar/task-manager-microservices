package taskdomain

import "errors"

var (
	ErrInvalidProjectId   = errors.New("invalid project id")
	ErrInvalidDescription = errors.New("invalid description")
)
