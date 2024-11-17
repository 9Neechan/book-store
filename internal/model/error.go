package model

import "github.com/pkg/errors"

var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorAuthorNotFound = errors.New("author not found")
	
)
