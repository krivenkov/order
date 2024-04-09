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

	GetItem(ctx context.Context, userID string, id string) (*Order, error)
	GetList(ctx context.Context, userID string, req *GetListRequest) ([]*Order, error)
	Count(ctx context.Context, userID string, req *GetCountRequest) (int, error)
}

type GetListRequest struct {
	IDs        option.Option[[]string]
	Q          option.Option[string]
	Orders     option.Option[[]*order.Order]
	Pagination option.Option[paginator.Pagination]
}

type GetCountRequest struct {
	IDs option.Option[[]string]
	Q   option.Option[string]
}
