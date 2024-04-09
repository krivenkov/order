package inner

import (
	"errors"

	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/order/pkg/api"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, model.ErrNotFound) {
		return api.ErrNotFound
	}

	if errors.Is(err, model.ErrMultiItems) {
		return api.ErrMultiItems
	}

	return status.Error(codes.Unknown, err.Error())
}

func fromOrders(source []*api.Order) []*order.Order {
	target := make([]*order.Order, 0, len(source))
	for _, s := range source {
		direction := "asc"

		if s.Direction == api.Direction_DESC {
			direction = "desc"
		}

		target = append(target, &order.Order{
			Column:    s.Column,
			Direction: direction,
		})
	}

	return target
}

func fromPagination(source *api.Pagination) *paginator.Pagination {
	if source == nil {
		return nil
	}

	return &paginator.Pagination{
		Limit:  int(source.Limit),
		Offset: int(source.Offset),
	}
}

func toOrderItem(source *orderModel.Order) *api.OrderItem {
	return &api.OrderItem{
		Id:          source.ID,
		Status:      api.OrderItemStatus(source.Status),
		TsCreate:    timestamppb.New(source.TSCreate),
		TsModify:    timestamppb.New(source.TSModify),
		UserId:      source.UserID,
		Name:        source.Name,
		Description: source.Description,
	}
}

func toOrderItemList(source []*orderModel.Order) []*api.OrderItem {
	target := make([]*api.OrderItem, 0, len(source))

	for _, s := range source {
		target = append(target, toOrderItem(s))
	}

	return target
}
