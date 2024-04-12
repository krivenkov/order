package list_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	orderModel "github.com/krivenkov/order/internal/model/order"
	orderMock "github.com/krivenkov/order/internal/model/order/mock"
	"github.com/krivenkov/order/internal/server/http/convertors"
	"github.com/krivenkov/order/internal/server/http/handlers/order/list"
	"github.com/krivenkov/order/internal/server/http/models"
	orderOperation "github.com/krivenkov/order/internal/server/http/operations/order"
	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
	"github.com/krivenkov/pkg/ptr"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := list.New(mock)

		var (
			userID      = "user_id"
			name        = "name"
			description = "description"
			limit       = 10
			offset      = 0
			i           interface{}
		)

		obj := &orderModel.Order{
			ID:          uuid.NewString(),
			TSCreate:    now(),
			TSModify:    now(),
			Status:      orderModel.StatusCreated,
			UserID:      userID,
			Name:        name,
			Description: description,
		}

		obj2 := &orderModel.Order{
			ID:          uuid.NewString(),
			TSCreate:    now(),
			TSModify:    now(),
			Status:      orderModel.StatusCreated,
			UserID:      userID,
			Name:        name + "2",
			Description: description,
		}

		serviceRes := []*orderModel.Order{obj, obj2}

		filter := &orderModel.GetListRequest{
			Orders: option.New([]*order.Order{
				{
					Column:    orderModel.NameSortKey,
					Direction: "desc",
				},
			}),
			Pagination: option.New(paginator.Pagination{
				Limit:  limit,
				Offset: offset,
			}),
		}

		filterCount := &orderModel.GetCountRequest{}

		mock.EXPECT().GetList(gomock.Any(), userID, filter).Return(serviceRes, nil)

		mock.EXPECT().Count(gomock.Any(), userID, filterCount).Return(2, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/order", nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrdersParams{
			HTTPRequest:   req,
			Limit:         ptr.Pointer(float64(limit)),
			Offset:        ptr.Pointer(float64(offset)),
			SortBy:        ptr.Pointer(orderModel.NameSortKey),
			SortDirection: ptr.Pointer("desc"),
		}, i)

		data := &models.GetOrdersResponse{
			Orders: convertors.OrdersFromModel(serviceRes),
			Pagination: convertors.Pagination(&paginator.PaginationResult{
				Limit:  limit,
				Offset: offset,
				Total:  2,
			}),
		}

		if respOk, ok := res.(*orderOperation.GetOrdersOK); ok {
			require.Equal(t, data, respOk.Payload)
		} else {
			require.FailNow(t, "resp is not GetOrdersOK")
		}

	})

	t.Run("Bad extract count", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := list.New(mock)

		var (
			userID      = "user_id"
			name        = "name"
			description = "description"
			limit       = 10
			offset      = 0
			i           interface{}
			someErr     = errors.New("some error")
		)

		obj := &orderModel.Order{
			ID:          uuid.NewString(),
			TSCreate:    now(),
			TSModify:    now(),
			Status:      orderModel.StatusCreated,
			UserID:      userID,
			Name:        name,
			Description: description,
		}

		obj2 := &orderModel.Order{
			ID:          uuid.NewString(),
			TSCreate:    now(),
			TSModify:    now(),
			Status:      orderModel.StatusCreated,
			UserID:      userID,
			Name:        name + "2",
			Description: description,
		}

		serviceRes := []*orderModel.Order{obj, obj2}

		filter := &orderModel.GetListRequest{
			Orders: option.New([]*order.Order{
				{
					Column:    orderModel.NameSortKey,
					Direction: "desc",
				},
			}),
			Pagination: option.New(paginator.Pagination{
				Limit:  limit,
				Offset: offset,
			}),
		}

		filterCount := &orderModel.GetCountRequest{}

		mock.EXPECT().GetList(gomock.Any(), userID, filter).Return(serviceRes, nil)

		mock.EXPECT().Count(gomock.Any(), userID, filterCount).Return(0, someErr)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/order", nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrdersParams{
			HTTPRequest:   req,
			Limit:         ptr.Pointer(float64(limit)),
			Offset:        ptr.Pointer(float64(offset)),
			SortBy:        ptr.Pointer(orderModel.NameSortKey),
			SortDirection: ptr.Pointer("desc"),
		}, i)

		require.Equal(t, orderOperation.NewGetOrdersInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Get list failed"),
		}), res)
	})

	t.Run("Bad extract list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := list.New(mock)

		var (
			userID  = "user_id"
			limit   = 10
			offset  = 0
			i       interface{}
			someErr = errors.New("some error")
		)

		filter := &orderModel.GetListRequest{
			Orders: option.New([]*order.Order{
				{
					Column:    orderModel.NameSortKey,
					Direction: "desc",
				},
			}),
			Pagination: option.New(paginator.Pagination{
				Limit:  limit,
				Offset: offset,
			}),
		}

		mock.EXPECT().GetList(gomock.Any(), userID, filter).Return(nil, someErr)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/order", nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrdersParams{
			HTTPRequest:   req,
			Limit:         ptr.Pointer(float64(limit)),
			Offset:        ptr.Pointer(float64(offset)),
			SortBy:        ptr.Pointer(orderModel.NameSortKey),
			SortDirection: ptr.Pointer("desc"),
		}, i)

		require.Equal(t, orderOperation.NewGetOrdersInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Get list failed"),
		}), res)
	})

	t.Run("Success full filters", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := list.New(mock)

		var (
			name        = "name"
			description = "description"
			userID      = "user_id"
			q           = "slug"
			limit       = 10
			offset      = 0
			i           interface{}
		)

		obj := &orderModel.Order{
			ID:          uuid.NewString(),
			TSCreate:    now(),
			TSModify:    now(),
			Status:      orderModel.StatusCreated,
			UserID:      userID,
			Name:        name,
			Description: description,
		}

		obj2 := &orderModel.Order{
			ID:          uuid.NewString(),
			TSCreate:    now(),
			TSModify:    now(),
			Status:      orderModel.StatusCreated,
			UserID:      userID,
			Name:        name + "2",
			Description: description,
		}

		serviceRes := []*orderModel.Order{obj, obj2}

		filter := &orderModel.GetListRequest{
			Q: option.New(q),
			Orders: option.New([]*order.Order{
				{
					Column:    orderModel.NameSortKey,
					Direction: "desc",
				},
			}),
			Pagination: option.New(paginator.Pagination{
				Limit:  limit,
				Offset: offset,
			}),
		}

		filterCount := &orderModel.GetCountRequest{
			Q: option.New(q),
		}

		mock.EXPECT().GetList(gomock.Any(), userID, filter).Return(serviceRes, nil)

		mock.EXPECT().Count(gomock.Any(), userID, filterCount).Return(2, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/order", nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrdersParams{
			HTTPRequest:   req,
			Limit:         ptr.Pointer(float64(limit)),
			Offset:        ptr.Pointer(float64(offset)),
			SortBy:        ptr.Pointer(orderModel.NameSortKey),
			SortDirection: ptr.Pointer("desc"),
			Q:             ptr.Pointer(q),
		}, i)

		data := &models.GetOrdersResponse{
			Orders: convertors.OrdersFromModel(serviceRes),
			Pagination: convertors.Pagination(&paginator.PaginationResult{
				Limit:  limit,
				Offset: offset,
				Total:  2,
			}),
		}

		if respOk, ok := res.(*orderOperation.GetOrdersOK); ok {
			require.Equal(t, data, respOk.Payload)
		} else {
			require.FailNow(t, "resp is not GetOrdersOK")
		}

	})
}

func now() time.Time {
	return time.Date(2000, time.January, 1, 15, 24, 11, 0, time.UTC)
}

func newID() uuid.UUID {
	return uuid.Nil
}
