package order

import (
	"context"
)

//go:generate mockgen -source=service.go -destination=mock/service.go

type Service interface {
	Create(ctx context.Context, userID string, form *Form) (*Order, error)
	Update(ctx context.Context, userID, id string, form *Form) (*Order, error)
	SoftDelete(ctx context.Context, userID, id string) error

	GetList(ctx context.Context, userID string, filter *Filter) ([]*Order, error)
	GetItem(ctx context.Context, userID string, id string) (*Order, error)
	Count(ctx context.Context, userID string, filter *Filter) (int, error)
}
