package remove_test

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
	orderMock "github.com/krivenkov/order/internal/model/order/mock"
	"github.com/krivenkov/order/internal/server/http/handlers/order/remove"
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
		serv := remove.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		mock.EXPECT().SoftDelete(gomock.Any(), userID, newID().String()).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.DeleteOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewDeleteOrderNoContent(), res)

	})

	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := remove.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		mock.EXPECT().SoftDelete(gomock.Any(), userID, newID().String()).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.DeleteOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewDeleteOrderNoContent(), res)

	})

	t.Run("Not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := remove.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		mock.EXPECT().SoftDelete(gomock.Any(), userID, newID().String()).Return(model.ErrNotFound)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.DeleteOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewDeleteOrderNotFound().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer(model.ErrNotFound.Error()),
		}), res)

	})

	t.Run("Permission denied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := remove.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		mock.EXPECT().SoftDelete(gomock.Any(), userID, newID().String()).Return(model.ErrPermissionDenied)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.DeleteOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewDeleteOrderForbidden().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorAccessDenied),
			ErrorDescription: ptr.Pointer(model.ErrPermissionDenied.Error()),
		}), res)

	})

	t.Run("Bad", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := remove.New(mock)

		var (
			userID = "user_id"
			i      interface{}

			someErr = errors.New("some error")
		)

		mock.EXPECT().SoftDelete(gomock.Any(), userID, newID().String()).Return(someErr)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/order/orders/%s", newID().String()), nil)
		i = userID

		res := serv.Handle(orderOperation.DeleteOrderParams{
			HTTPRequest: req,
			ID:          newID().String(),
		}, i)

		require.Equal(t, orderOperation.NewDeleteOrderInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Delete order failed"),
		}), res)

	})

	t.Run("Bad id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		mock := orderMock.NewMockService(ctrl)
		serv := remove.New(mock)

		var (
			userID = "user_id"
			i      interface{}
		)

		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/order/orders/%s", "123"), nil)
		i = userID

		res := serv.Handle(orderOperation.DeleteOrderParams{
			HTTPRequest: req,
			ID:          "123",
		}, i)

		require.Equal(t, orderOperation.NewDeleteOrderNotFound().WithPayload(&models.Error{
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
