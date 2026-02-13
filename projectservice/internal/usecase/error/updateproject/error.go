package updateerr

import "errors"

var (
	ErrProjectNotFound          = errors.New("project not found")
	ErrProjectNameAlreadyExists = errors.New("project with this name already exists")
)
