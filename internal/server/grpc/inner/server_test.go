package inner_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	orderMock "github.com/krivenkov/order/internal/model/order/mock"
	"github.com/krivenkov/order/internal/server/grpc/inner"
	"github.com/krivenkov/order/pkg/api"
	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetOrderItem(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			svc = orderMock.NewMockService(ctrl)
		)

		svc.EXPECT().InnerGetItem(context.TODO(), &orderModel.InnerGetItemRequest{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
		}).Return(orderItem, nil)

		srv := inner.NewServer(svc)

		res, err := srv.GetOrderItem(context.TODO(), &api.OrderItemRequest{
			Filter: &api.OrderItemFilter{
				Ids:    []string{newID().String()},
				UserId: &userID,
			},
		})

		require.NoError(t, err)
		require.Equal(t, &api.OrderItemResponse{
			Value: &api.OrderItem{
				Id:          newID().String(),
				Status:      api.OrderItemStatus_StatusCreated,
				TsCreate:    timestamppb.New(now()),
				TsModify:    timestamppb.New(now()),
				UserId:      userID,
				Name:        name,
				Description: description,
			},
		}, res)
	})

	t.Run("Bad", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"

			svc = orderMock.NewMockService(ctrl)
		)

		svc.EXPECT().InnerGetItem(context.TODO(), &orderModel.InnerGetItemRequest{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
		}).Return(nil, model.ErrNotFound)

		srv := inner.NewServer(svc)

		res, err := srv.GetOrderItem(context.TODO(), &api.OrderItemRequest{
			Filter: &api.OrderItemFilter{
				Ids:    []string{newID().String()},
				UserId: &userID,
			},
		})

		require.ErrorIs(t, err, api.ErrNotFound)
		require.Nil(t, res)
	})
}

func TestGetOrderItemList(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			orders = []*orderModel.Order{orderItem}

			svc = orderMock.NewMockService(ctrl)
		)

		svc.EXPECT().InnerGetList(context.TODO(), &orderModel.InnerGetListRequest{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
			Orders: option.New([]*order.Order{
				{
					Column:    orderModel.NameSortKey,
					Direction: "asc",
				},
			}),
			Pagination: option.New(paginator.Pagination{
				Limit:  10,
				Offset: 0,
			}),
		}).Return(orders, nil)

		srv := inner.NewServer(svc)

		res, err := srv.GetOrderItemList(context.TODO(), &api.OrderItemListRequest{
			Filter: &api.OrderItemFilter{
				Ids:    []string{newID().String()},
				UserId: &userID,
			},
			Orders: []*api.Order{
				{
					Column:    orderModel.NameSortKey,
					Direction: api.Direction_ASC,
				},
			},
			Pagination: &api.Pagination{
				Limit:  10,
				Offset: 0,
			},
		})

		require.NoError(t, err)
		require.Equal(t, &api.OrderItemListResponse{
			Value: []*api.OrderItem{{
				Id:          newID().String(),
				Status:      api.OrderItemStatus_StatusCreated,
				TsCreate:    timestamppb.New(now()),
				TsModify:    timestamppb.New(now()),
				UserId:      userID,
				Name:        name,
				Description: description,
			}},
		}, res)
	})

	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"

			someErr = fmt.Errorf("some error")

			svc = orderMock.NewMockService(ctrl)
		)

		svc.EXPECT().InnerGetList(context.TODO(), &orderModel.InnerGetListRequest{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
			Orders: option.New([]*order.Order{
				{
					Column:    orderModel.NameSortKey,
					Direction: "asc",
				},
			}),
			Pagination: option.New(paginator.Pagination{
				Limit:  10,
				Offset: 0,
			}),
		}).Return(nil, someErr)

		srv := inner.NewServer(svc)

		res, err := srv.GetOrderItemList(context.TODO(), &api.OrderItemListRequest{
			Filter: &api.OrderItemFilter{
				Ids:    []string{newID().String()},
				UserId: &userID,
			},
			Orders: []*api.Order{
				{
					Column:    orderModel.NameSortKey,
					Direction: api.Direction_ASC,
				},
			},
			Pagination: &api.Pagination{
				Limit:  10,
				Offset: 0,
			},
		})

		require.Error(t, err)
		require.Nil(t, res)
	})
}

func now() time.Time {
	return time.Date(2000, time.January, 1, 15, 24, 11, 0, time.UTC)
}

func newID() uuid.UUID {
	return uuid.Nil
}
