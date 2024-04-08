package order

import (
	"context"

	"gitlab.bred.team/bre-back/pkg/option"
	"gitlab.bred.team/bre-back/pkg/order"
	"gitlab.bred.team/bre-back/pkg/paginator"
)

//go:generate mockgen -source=querier.go -destination=mock/querier.go

type Querier interface {
	GetItem(ctx context.Context, filter *Filter) (*Order, error)
	GetList(ctx context.Context, filter *Filter, orders option.Option[[]*order.Order], pagination option.Option[paginator.Pagination]) ([]*Order, error)
	Count(ctx context.Context, filter *Filter) (int, error)
}

type Filter struct {
	IDs    option.Option[[]string]
	Status option.Option[int]
	UserID option.Option[string]
}
