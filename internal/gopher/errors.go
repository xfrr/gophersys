package gopher

import "errors"

var (
	// ErrInvalidID represents an error when the ID is invalid
	ErrInvalidID = errors.New("invalid id")

	// ErrInvalidStatus represents an error when the status is invalid
	ErrInvalidStatus = errors.New("invalid status")

	// ErrInvalidUsername represents an error when the username is invalid
	ErrInvalidUsername = errors.New("invalid username")

	// ErrNotFound represents an error when the gopher is not found
	ErrNotFound = errors.New("gopher not found")

	// ErrAlreadyExistsByUsernameOrID represents an error when the gopher already exists
	ErrAlreadyExistsByUsernameOrID = errors.New("gopher already exists with given username or id")
)
