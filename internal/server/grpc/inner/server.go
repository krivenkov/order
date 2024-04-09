package inner

import (
	"context"

	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/order/pkg/api"
	"github.com/krivenkov/pkg/option"
)

type server struct {
	svc orderModel.Service
}

func NewServer(svc orderModel.Service) api.OrderServiceServer {
	return &server{
		svc: svc,
	}
}

func (s *server) GetOrderItem(ctx context.Context, request *api.OrderItemRequest) (*api.OrderItemResponse, error) {
	filter := &orderModel.InnerGetItemRequest{}

	if request.Filter != nil {
		if len(request.Filter.Ids) > 0 {
			filter.IDs = option.New(request.Filter.Ids)
		}

		if request.Filter.UserId != nil {
			filter.UserID = option.New(*request.Filter.UserId)
		}
	}

	item, err := s.svc.InnerGetItem(ctx, filter)
	if err != nil {
		return nil, toError(err)
	}

	return &api.OrderItemResponse{
		Value: toOrderItem(item),
	}, nil
}

func (s *server) GetOrderItemList(ctx context.Context, request *api.OrderItemListRequest) (*api.OrderItemListResponse, error) {
	var (
		filter     = &orderModel.InnerGetListRequest{}
		orders     = fromOrders(request.Orders)
		pagination = fromPagination(request.Pagination)
	)

	if request.Filter != nil {
		if len(request.Filter.Ids) > 0 {
			filter.IDs = option.New(request.Filter.Ids)
		}

		if request.Filter.UserId != nil {
			filter.UserID = option.New(*request.Filter.UserId)
		}

		if len(orders) > 0 {
			filter.Orders = option.New(orders)
		}

		if pagination != nil {
			filter.Pagination = option.New(*pagination)
		}
	}

	items, err := s.svc.InnerGetList(ctx, filter)
	if err != nil {
		return nil, toError(err)
	}

	return &api.OrderItemListResponse{
		Value: toOrderItemList(items),
	}, nil
}
