package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("User doesn't exists")
	ErrCollectioNotFound = errors.New("Collection doesn't exists")
)
