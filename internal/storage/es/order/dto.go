package order

import (
	"github.com/krivenkov/order/internal/model/order"
)

const (
	indexName = "order"
)

type dto struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Status      int64  `json:"status"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func newDto() *dto {
	return &dto{}
}

func (d *dto) toModel() *order.Order {
	if d == nil {
		return nil
	}

	return &order.Order{
		ID:          d.ID,
		Status:      order.Status(d.Status),
		Name:        d.Name,
		Description: d.Description,
	}
}

func (d *dto) fromModel(source *order.Order) {
	target := dto{
		ID:          source.ID,
		Status:      int64(source.Status),
		UserID:      source.UserID,
		Name:        source.Name,
		Description: source.Description,
	}

	*d = target
}
