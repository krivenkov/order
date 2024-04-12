package order_test

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
	svc "github.com/krivenkov/order/internal/service/order"
	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
	txerMock "github.com/krivenkov/pkg/txer/mock"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGCommander.EXPECT().Create(context.TODO(), orderItem).Return(nil)

		orderESCommander.EXPECT().Create(context.TODO(), orderItem).Return(nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Create(context.TODO(), userID, &orderModel.Form{
			Name:        &name,
			Description: &description,
		})

		require.NoError(t, err)
		require.Equal(t, orderItem, res)
	})

	t.Run("Error save in es", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGCommander.EXPECT().Create(context.TODO(), orderItem).Return(nil)

		orderESCommander.EXPECT().Create(context.TODO(), orderItem).Return(someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Create(context.TODO(), userID, &orderModel.Form{
			Name:        &name,
			Description: &description,
		})

		require.ErrorIs(t, err, someErr)
		require.Nil(t, res)
	})

	t.Run("Error save in pg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGCommander.EXPECT().Create(context.TODO(), orderItem).Return(someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Create(context.TODO(), userID, &orderModel.Form{
			Name:        &name,
			Description: &description,
		})

		require.ErrorIs(t, err, someErr)
		require.Nil(t, res)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		orderPGCommander.EXPECT().Update(context.TODO(), orderItem).Return(nil)

		orderESCommander.EXPECT().Update(context.TODO(), orderItem).Return(nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Update(context.TODO(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		})

		require.NoError(t, err)
		require.Equal(t, orderItem, res)
	})

	t.Run("Error update in es", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		orderPGCommander.EXPECT().Update(context.TODO(), orderItem).Return(nil)

		orderESCommander.EXPECT().Update(context.TODO(), orderItem).Return(someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Update(context.TODO(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		})

		require.ErrorIs(t, err, someErr)
		require.Nil(t, res)
	})

	t.Run("Error update in pg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		orderPGCommander.EXPECT().Update(context.TODO(), orderItem).Return(someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Update(context.TODO(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		})

		require.ErrorIs(t, err, someErr)
		require.Nil(t, res)
	})

	t.Run("Error get item in pg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Update(context.TODO(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		})

		require.ErrorIs(t, err, someErr)
		require.Nil(t, res)
	})

	t.Run("Error permission", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID + "bad",
				Name:        name,
				Description: description,
			}
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Update(context.TODO(), userID, newID().String(), &orderModel.Form{
			Name:        &name,
			Description: &description,
		})

		require.ErrorIs(t, err, model.ErrPermissionDenied)
		require.Nil(t, res)
	})
}

func TestSoftDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		orderPGCommander.EXPECT().Update(context.TODO(), orderItem).Return(nil)

		orderESCommander.EXPECT().Update(context.TODO(), orderItem).Return(nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		err := service.SoftDelete(context.TODO(), userID, newID().String())

		require.NoError(t, err)
	})

	t.Run("Error update in es", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		orderPGCommander.EXPECT().Update(context.TODO(), orderItem).Return(nil)

		orderESCommander.EXPECT().Update(context.TODO(), orderItem).Return(someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		err := service.SoftDelete(context.TODO(), userID, newID().String())

		require.ErrorIs(t, err, someErr)
	})

	t.Run("Error update in pg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		orderPGCommander.EXPECT().Update(context.TODO(), orderItem).Return(someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		err := service.SoftDelete(context.TODO(), userID, newID().String())

		require.ErrorIs(t, err, someErr)
	})

	t.Run("Error get item in pg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		err := service.SoftDelete(context.TODO(), userID, newID().String())

		require.ErrorIs(t, err, someErr)
	})

	t.Run("Error permission", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID + "bad",
				Name:        name,
				Description: description,
			}
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		err := service.SoftDelete(context.TODO(), userID, newID().String())

		require.ErrorIs(t, err, model.ErrPermissionDenied)
	})
}

func TestDisable(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGCommander.EXPECT().Disable(context.TODO(), userID).Return(nil)

		orderESCommander.EXPECT().Disable(context.TODO(), userID).Return(nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		err := service.Disable(context.TODO(), userID)

		require.NoError(t, err)
	})

	t.Run("Bad save in es", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGCommander.EXPECT().Disable(context.TODO(), userID).Return(nil)

		orderESCommander.EXPECT().Disable(context.TODO(), userID).Return(someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		err := service.Disable(context.TODO(), userID)

		require.ErrorIs(t, err, someErr)
	})

	t.Run("Bad save in pg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			someErr = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGCommander.EXPECT().Disable(context.TODO(), userID).Return(someErr)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		err := service.Disable(context.TODO(), userID)

		require.ErrorIs(t, err, someErr)
	})
}

func TestGetItem(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.GetItem(context.TODO(), userID, newID().String())

		require.NoError(t, err)
		require.Equal(t, res, orderItem)
	})

	t.Run("Permission denied", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID + "bad",
				Name:        name,
				Description: description,
			}
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(orderItem, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.GetItem(context.TODO(), userID, newID().String())

		require.ErrorIs(t, err, model.ErrPermissionDenied)
		require.Nil(t, res)
	})

	t.Run("Bad", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			someError = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(nil, someError)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.GetItem(context.TODO(), userID, newID().String())

		require.ErrorIs(t, err, someError)
		require.Nil(t, res)
	})
}

func TestInnerGetItem(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			orderItem = &orderModel.Order{
				ID:          newID().String(),
				TSCreate:    now(),
				TSModify:    now(),
				Status:      orderModel.StatusCreated,
				UserID:      userID,
				Name:        name,
				Description: description,
			}
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
		}).Return(orderItem, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.InnerGetItem(context.TODO(), &orderModel.InnerGetItemRequest{
			UserID: option.New(userID),
			IDs:    option.New([]string{newID().String()}),
		})

		require.NoError(t, err)
		require.Equal(t, res, orderItem)
	})

	t.Run("Bad", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

			someError = fmt.Errorf("some error")
		)

		tXer.EXPECT().WithTX(context.TODO(), gomock.Any()).DoAndReturn(func(ctx context.Context, cb func(ctx context.Context) error) error {
			return cb(ctx)
		}).AnyTimes()

		orderPGQuerier.EXPECT().GetItem(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
		}).Return(nil, someError)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.InnerGetItem(context.TODO(), &orderModel.InnerGetItemRequest{
			UserID: option.New(userID),
			IDs:    option.New([]string{newID().String()}),
		})

		require.ErrorIs(t, err, someError)
		require.Nil(t, res)
	})
}

func TestCount(t *testing.T) {
	t.Run("Success in pg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"
			count  = 103

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)
		)

		orderPGQuerier.EXPECT().Count(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(count, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Count(context.TODO(), userID, &orderModel.GetCountRequest{
			IDs: option.New([]string{newID().String()}),
		})

		require.NoError(t, err)
		require.Equal(t, res, count)
	})

	t.Run("Success in es", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"
			q      = "test"
			count  = 103

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)
		)

		orderESQuerier.EXPECT().Count(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
			Status: option.New(int(orderModel.StatusCreated)),
			Q:      option.New(q),
		}).Return(count, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Count(context.TODO(), userID, &orderModel.GetCountRequest{
			Q:   option.New(q),
			IDs: option.New([]string{newID().String()}),
		})

		require.NoError(t, err)
		require.Equal(t, res, count)
	})

	t.Run("Success empty filter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID = "user_id"
			count  = 103

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)
		)

		orderPGQuerier.EXPECT().Count(context.TODO(), &orderModel.Filter{
			UserID: option.New(userID),
			Status: option.New(int(orderModel.StatusCreated)),
		}).Return(count, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.Count(context.TODO(), userID, nil)

		require.NoError(t, err)
		require.Equal(t, res, count)
	})
}

func TestGetList(t *testing.T) {
	t.Run("Success in pg", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

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
		)

		orderPGQuerier.EXPECT().GetList(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
			Status: option.New(int(orderModel.StatusCreated)),
		}, []*order.Order{
			{
				Column:    orderModel.NameSortKey,
				Direction: "asc",
			},
		}, &paginator.Pagination{
			Limit:  10,
			Offset: 0,
		}).Return(orders, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.GetList(context.TODO(), userID, &orderModel.GetListRequest{
			IDs: option.New([]string{newID().String()}),
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
		})

		require.NoError(t, err)
		require.Equal(t, res, orders)
	})

	t.Run("Success in es", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			q           = "test"
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

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
		)

		orderESQuerier.EXPECT().GetList(context.TODO(), &orderModel.Filter{
			Q:      option.New(q),
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
			Status: option.New(int(orderModel.StatusCreated)),
		}, []*order.Order{
			{
				Column:    orderModel.NameSortKey,
				Direction: "asc",
			},
		}, &paginator.Pagination{
			Limit:  10,
			Offset: 0,
		}).Return(orders, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.GetList(context.TODO(), userID, &orderModel.GetListRequest{
			Q:   option.New(q),
			IDs: option.New([]string{newID().String()}),
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
		})

		require.NoError(t, err)
		require.Equal(t, res, orders)
	})
}

func TestInnerGetList(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var (
			userID      = "user_id"
			name        = "test"
			description = "some text"

			orderPGQuerier   = orderMock.NewMockQuerier(ctrl)
			orderESQuerier   = orderMock.NewMockQuerier(ctrl)
			orderPGCommander = orderMock.NewMockCommander(ctrl)
			orderESCommander = orderMock.NewMockCommander(ctrl)
			tXer             = txerMock.NewMockTXer(ctrl)

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
		)

		orderPGQuerier.EXPECT().GetList(context.TODO(), &orderModel.Filter{
			IDs:    option.New([]string{newID().String()}),
			UserID: option.New(userID),
		}, []*order.Order{
			{
				Column:    orderModel.NameSortKey,
				Direction: "asc",
			},
		}, &paginator.Pagination{
			Limit:  10,
			Offset: 0,
		}).Return(orders, nil)

		service := svc.New(svc.Params{
			CmdPg: orderPGCommander,
			CmdEs: orderESCommander,
			QrPg:  orderPGQuerier,
			QrEs:  orderESQuerier,
			TXer:  tXer,
			Now:   now,
			NewID: newID,
		})

		res, err := service.InnerGetList(context.TODO(), &orderModel.InnerGetListRequest{
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
		})

		require.NoError(t, err)
		require.Equal(t, res, orders)
	})
}

func now() time.Time {
	return time.Date(2000, time.January, 1, 15, 24, 11, 0, time.UTC)
}

func newID() uuid.UUID {
	return uuid.Nil
}
