package domain

import "errors"

var (
	// ErrDataNotFound is an error for when requested data is not found
	ErrDataNotFound = errors.New("data not found")
	// ErrNoCreatedData is an error for when no data is provided to update
	ErrNoCreatedData = errors.New("no data was created")
)
