package convertors

import (
	"github.com/go-openapi/strfmt"
	"github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/order/internal/server/http/models"
	"github.com/krivenkov/pkg/ptr"
)

func OrderFromModel(n *order.Order) *models.Order {
	if n == nil {
		return nil
	}

	return &models.Order{
		ID:          ptr.Pointer(strfmt.UUID(n.ID)),
		Name:        ptr.Pointer(n.Name),
		Description: ptr.Pointer(n.Description),
	}
}

func OrdersFromModel(items []*order.Order) []*models.Order {
	orders := make([]*models.Order, 0, len(items))

	for _, item := range items {
		orders = append(orders, OrderFromModel(item))
	}
	return orders
}
