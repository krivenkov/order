package create_test

import (
	"bytes"
	"encoding/json"
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
	"github.com/krivenkov/order/internal/server/http/handlers/order/create"
	"github.com/krivenkov/order/internal/server/http/models"
	"github.com/krivenkov/order/internal/server/http/operations/order"
	"github.com/krivenkov/pkg/ptr"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := create.New(mock)

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

		mock.EXPECT().Create(gomock.Any(), userID, &orderModel.Form{
			Name:        &name,
			Description: &description,
		}).Return(obj, nil)

		reqBody := &models.CreateOrderRequest{
			Name:        &name,
			Description: &description,
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/order", bytes.NewReader(body))
		i = userID

		res := serv.Handle(order.CreateOrderParams{
			HTTPRequest: req,
			Body:        reqBody,
		}, i)

		data := &models.CreateOrderResponse{
			Order: convertors.OrderFromModel(obj),
		}

		if respOk, ok := res.(*order.CreateOrderOK); ok {
			require.Equal(t, data, respOk.Payload)
		} else {
			require.FailNow(t, "resp is not CreateNotesOK")
		}

	})

	t.Run("Bad", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := create.New(mock)

		var (
			userID      = "user_id"
			i           interface{}
			name        = "name"
			description = "description"

			someErr = errors.New("some error")
		)

		mock.EXPECT().Create(gomock.Any(), userID, &orderModel.Form{
			Name:        &name,
			Description: &description,
		}).Return(nil, someErr)

		reqBody := &models.CreateOrderRequest{
			Name:        &name,
			Description: &description,
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/order", bytes.NewReader(body))
		i = userID

		res := serv.Handle(order.CreateOrderParams{
			HTTPRequest: req,
			Body:        reqBody,
		}, i)

		require.Equal(t, order.NewCreateOrderInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Create order failed"),
		}), res)
	})

	t.Run("Empty body", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := create.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/order", nil)
		i = userID

		res := serv.Handle(order.CreateOrderParams{
			HTTPRequest: req,
		}, i)

		require.Equal(t, order.NewCreateOrderBadRequest().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("request body is empty"),
		}), res)
	})
}

func now() time.Time {
	return time.Date(2000, time.January, 1, 15, 24, 11, 0, time.UTC)
}

func newID() uuid.UUID {
	return uuid.Nil
}
