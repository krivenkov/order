package order

import (
	"context"
)

//go:generate mockgen -source=commander.go -destination=mock/commander.go

type Commander interface {
	Create(ctx context.Context, item *Order) error
	Update(ctx context.Context, item *Order) error
	Delete(ctx context.Context, item *Order) error
	Disable(ctx context.Context, userID string) error
}
