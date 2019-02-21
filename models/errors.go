package models

import "errors"

var (
	// ErrPostNotFound caises when post not found in database
	ErrPostNotFound = errors.New("requested post not found")
)
