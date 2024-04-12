package order

import (
	"context"

	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
)

//go:generate mockgen -source=querier.go -destination=mock/querier.go

type Querier interface {
	GetItem(ctx context.Context, filter *Filter) (*Order, error)
	GetList(ctx context.Context, filter *Filter, orders []*order.Order, pagination *paginator.Pagination) ([]*Order, error)
	Count(ctx context.Context, filter *Filter) (int, error)
}

type Filter struct {
	IDs    option.Option[[]string]
	Status option.Option[int]
	UserID option.Option[string]
	Q      option.Option[string]
}
