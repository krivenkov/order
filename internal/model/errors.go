package model

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrMultiItems = errors.New("multi items")
)
