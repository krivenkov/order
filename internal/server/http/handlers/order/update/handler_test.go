package update_test

import (
	"bytes"
	"encoding/json"
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
	"github.com/krivenkov/order/internal/server/http/handlers/order/update"
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
		serv := update.New(mock)

		var (
			name        = "name"
			description = "description"
			userID      = "user_id"
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

		mock.EXPECT().Update(gomock.Any(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		}).Return(obj, nil)

		reqBody := &models.UpdateOrderRequest{
			Name:        &name,
			Description: &description,
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), bytes.NewReader(body))
		i = userID

		res := serv.Handle(orderOperation.UpdateOrderParams{
			HTTPRequest: req,
			Body:        reqBody,
			ID:          newID().String(),
		}, i)

		data := &models.UpdateOrderResponse{
			Order: convertors.OrderFromModel(obj),
		}

		if respOk, ok := res.(*orderOperation.UpdateOrderOK); ok {
			require.Equal(t, data, respOk.Payload)
		} else {
			require.FailNow(t, "resp is not UpdateOrdersOK")
		}

	})

	t.Run("Not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := update.New(mock)

		var (
			name        = "name"
			description = "description"
			userID      = "user_id"
			i           interface{}
		)

		mock.EXPECT().Update(gomock.Any(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		}).Return(nil, model.ErrNotFound)

		reqBody := &models.UpdateOrderRequest{
			Name:        &name,
			Description: &description,
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/order/orders//%s", newID().String()), bytes.NewReader(body))
		i = userID

		res := serv.Handle(orderOperation.UpdateOrderParams{
			HTTPRequest: req,
			Body:        reqBody,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewUpdateOrderNotFound().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer(model.ErrNotFound.Error()),
		}), res)
	})

	t.Run("Permission denied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := update.New(mock)

		var (
			name        = "name"
			description = "description"
			userID      = "user_id"
			i           interface{}
		)

		mock.EXPECT().Update(gomock.Any(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		}).Return(nil, model.ErrPermissionDenied)

		reqBody := &models.UpdateOrderRequest{
			Name:        &name,
			Description: &description,
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), bytes.NewReader(body))
		i = userID

		res := serv.Handle(orderOperation.UpdateOrderParams{
			HTTPRequest: req,
			Body:        reqBody,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewUpdateOrderForbidden().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorAccessDenied),
			ErrorDescription: ptr.Pointer(model.ErrPermissionDenied.Error()),
		}), res)
	})

	t.Run("Some error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := update.New(mock)

		var (
			name        = "name"
			description = "description"
			userID      = "user_id"
			i           interface{}

			someErr = errors.New("some error")
		)

		mock.EXPECT().Update(gomock.Any(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		}).Return(nil, someErr)

		reqBody := &models.UpdateOrderRequest{
			Name:        &name,
			Description: &description,
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), bytes.NewReader(body))
		i = userID

		res := serv.Handle(orderOperation.UpdateOrderParams{
			HTTPRequest: req,
			Body:        reqBody,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewUpdateOrderInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Update order failed"),
		}), res)
	})

	t.Run("Empty body", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := update.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.UpdateOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewUpdateOrderBadRequest().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("request body is empty"),
		}), res)
	})

	t.Run("Bad id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := update.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/order/orders/%s", "123"), nil)
		i = userID

		res := serv.Handle(orderOperation.UpdateOrderParams{
			HTTPRequest: req,
			ID:          "123",
		}, i)

		require.Equal(t, orderOperation.NewUpdateOrderNotFound().WithPayload(&models.Error{
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
