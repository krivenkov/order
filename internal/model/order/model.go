package order

import "time"

type Status int

const (
	StatusCreated Status = 1
	StatusDeleted Status = 2
	StatusPayed   Status = 3
)

type Order struct {
	ID       string
	TSCreate time.Time
	TSModify time.Time
	Status   Status

	UserID string

	Name        string
	Description string
}
