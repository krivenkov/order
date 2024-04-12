package item_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	orderMock "github.com/krivenkov/order/internal/model/order/mock"
	"github.com/krivenkov/order/internal/server/http/convertors"
	"github.com/krivenkov/order/internal/server/http/handlers/order/item"
	"github.com/krivenkov/order/internal/server/http/models"
	orderOperation "github.com/krivenkov/order/internal/server/http/operations/order"
	"github.com/krivenkov/pkg/ptr"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := item.New(mock)

		var (
			userID      = "user_id"
			i           interface{}
			name        = "name"
			description = "description"
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

		mock.EXPECT().GetItem(gomock.Any(), userID, newID().String()).Return(obj, nil)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		data := &models.GetOrderResponse{
			Order: convertors.OrderFromModel(obj),
		}

		if respOk, ok := res.(*orderOperation.GetOrderOK); ok {
			require.Equal(t, data, respOk.Payload)
		} else {
			require.FailNow(t, "resp is not GetOrderOK")
		}

	})

	t.Run("Not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := item.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		mock.EXPECT().GetItem(gomock.Any(), userID, newID().String()).Return(nil, model.ErrNotFound)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewGetOrderNotFound().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer(model.ErrNotFound.Error()),
		}), res)

	})

	t.Run("Permission denied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := item.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		mock.EXPECT().GetItem(gomock.Any(), userID, newID().String()).Return(nil, model.ErrPermissionDenied)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewGetOrderForbidden().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorAccessDenied),
			ErrorDescription: ptr.Pointer(model.ErrPermissionDenied.Error()),
		}), res)

	})

	t.Run("Bad", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := item.New(mock)

		var (
			userID  = "user_id"
			i       interface{}
			someErr = errors.New("some error")
		)

		mock.EXPECT().GetItem(gomock.Any(), userID, newID().String()).Return(nil, someErr)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewGetOrderInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Get order failed"),
		}), res)

	})

	t.Run("Invalid id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := item.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.GetOrderParams{
			HTTPRequest: req,
			ID:          "123",
		}, i)

		require.Equal(t, orderOperation.NewGetOrderNotFound().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("Not Found"),
		}), res)

	})
}

func now() time.Time {
	return time.Date(2000, time.January, 1, 15, 24, 11, 0, time.UTC)
}

func newID() uuid.UUID {
	return uuid.Nil
}
