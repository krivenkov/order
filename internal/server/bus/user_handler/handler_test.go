package user_handler_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	orderMock "github.com/krivenkov/order/internal/model/order/mock"
	"github.com/krivenkov/order/internal/model/user"
	"github.com/krivenkov/order/internal/server/bus/user_handler"
	"github.com/krivenkov/pkg/bus"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		svc = orderMock.NewMockService(ctrl)

		userID = uuid.New().String()

		someErr = fmt.Errorf("some error")
	)

	handler := user_handler.New(svc)

	t.Run("NoPayload", func(t *testing.T) {
		res := handler.Handle(context.TODO(), bus.Message[user.User]{})
		require.Equal(t, bus.StatusWarning, res.Code)
	})

	t.Run("ActiveUser", func(t *testing.T) {
		res := handler.Handle(context.TODO(), bus.Message[user.User]{
			Value: bus.MessageValue[user.User]{
				Payload: &user.User{
					ID:      userID,
					Disable: false,
				},
			},
		})
		require.Equal(t, bus.StatusOk, res.Code)
	})

	t.Run("Success", func(t *testing.T) {
		svc.EXPECT().Disable(context.TODO(), userID).Return(nil)

		res := handler.Handle(context.TODO(), bus.Message[user.User]{
			Value: bus.MessageValue[user.User]{
				Payload: &user.User{
					ID:      userID,
					Disable: true,
				},
			},
		})
		require.Equal(t, bus.StatusOk, res.Code)
	})

	t.Run("Bad", func(t *testing.T) {
		svc.EXPECT().Disable(context.TODO(), userID).Return(someErr)

		res := handler.Handle(context.TODO(), bus.Message[user.User]{
			Value: bus.MessageValue[user.User]{
				Payload: &user.User{
					ID:      userID,
					Disable: true,
				},
			},
		})
		require.Equal(t, bus.StatusError, res.Code)
	})
}
