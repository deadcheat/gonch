package gonch

import "errors"

var (
	// ErrNotExists file does not exist
	ErrNotExists = errors.New("file does not exist")
	// ErrInvalidStatus file is not creatable
	ErrInvalidStatus = errors.New("file is invalid status")
)
