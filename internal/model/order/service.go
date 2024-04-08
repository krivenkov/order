package order

import (
	"context"

	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
)

//go:generate mockgen -source=service.go -destination=mock/service.go

type Service interface {
	Create(ctx context.Context, userID string, form *Form) (*Order, error)
	Update(ctx context.Context, userID, id string, form *Form) (*Order, error)
	SoftDelete(ctx context.Context, userID, id string) error

	GetList(ctx context.Context, userID string, filter option.Option[Filter], orders option.Option[[]*order.Order], pagination option.Option[paginator.Pagination]) ([]*Order, error)
	GetItem(ctx context.Context, userID string, id string) (*Order, error)
	Count(ctx context.Context, userID string, filter option.Option[Filter]) (int, error)
}
