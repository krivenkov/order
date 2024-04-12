package count_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	orderModel "github.com/krivenkov/order/internal/model/order"
	orderMock "github.com/krivenkov/order/internal/model/order/mock"
	"github.com/krivenkov/order/internal/server/http/handlers/order/count"
	"github.com/krivenkov/order/internal/server/http/models"
	"github.com/krivenkov/order/internal/server/http/operations/order"
	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/ptr"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := count.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		filter := &orderModel.GetCountRequest{}

		mock.EXPECT().Count(gomock.Any(), userID, filter).Return(2, nil)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/order/orders/count", nil)
		i = userID

		res := serv.Handle(order.GetOrdersCountParams{
			HTTPRequest: req,
		}, i)

		data := &models.GetCountResponse{
			Count: ptr.Pointer(int64(2)),
		}

		if respOk, ok := res.(*order.GetOrdersCountOK); ok {
			require.Equal(t, data, respOk.Payload)
		} else {
			require.FailNow(t, "resp is not GetCountOK")
		}

	})

	t.Run("Bad", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := count.New(mock)

		var (
			q       = "test"
			userID  = "user_id"
			i       interface{}
			someErr = errors.New("some error")
		)

		filter := &orderModel.GetCountRequest{
			Q: option.New(q),
		}

		mock.EXPECT().Count(gomock.Any(), userID, filter).Return(0, someErr)

		req := httptest.NewRequest(http.MethodGet, "/api/v1/order/orders/count", nil)
		i = userID

		res := serv.Handle(order.GetOrdersCountParams{
			Q:           &q,
			HTTPRequest: req,
		}, i)

		require.Equal(t, order.NewGetOrdersCountInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Get order count failed"),
		}), res)
	})
}
