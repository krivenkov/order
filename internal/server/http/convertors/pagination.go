package convertors

import (
	"github.com/krivenkov/order/internal/server/http/models"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
	"github.com/krivenkov/pkg/ptr"
)

func Pagination(s *paginator.PaginationResult) *models.Pagination {
	if s == nil {
		return nil
	}

	return &models.Pagination{
		Limit:  ptr.Pointer(float64(s.Limit)),
		Offset: ptr.Pointer(float64(s.Offset)),
		Total:  ptr.Pointer(float64(s.Total)),
	}
}

func Paginator(limit, offset *float64) *paginator.Pagination {
	res := &paginator.Pagination{}

	if limit != nil {
		res.Limit = int(*limit)
	}

	if offset != nil {
		res.Offset = int(*offset)
	}

	return res
}

func Order(sortBy, sortDirection *string) *order.Order {
	res := &order.Order{}

	if sortBy != nil {
		res.Column = *sortBy
	}

	if sortDirection != nil {
		res.Direction = *sortDirection
	}

	return res
}
