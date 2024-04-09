package order

import (
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	StatusCreated Status = 1
	StatusDeleted Status = 2
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

func New(userID string, now func() time.Time, newID func() uuid.UUID) *Order {
	return &Order{
		ID:       newID().String(),
		TSCreate: now(),
		TSModify: now(),
		Status:   StatusCreated,
		UserID:   userID,
	}
}

func (o *Order) FillForm(f *Form) {
	if f.Name != nil {
		o.Name = *f.Name
	}

	if f.Description != nil {
		o.Name = *f.Description
	}
}

type Form struct {
	Name        *string
	Description *string
}
