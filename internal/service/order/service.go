package order

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/txer"
	"go.uber.org/fx"
)

type service struct {
	cmdPg, cmdEs orderModel.Commander
	qrPg, qrEs   orderModel.Querier

	tXer  txer.TXer
	now   func() time.Time
	newID func() uuid.UUID
}

type Params struct {
	fx.In

	CmdPg orderModel.Commander `name:"order_pg_cmd"`
	CmdEs orderModel.Commander `name:"order_es_cmd"`
	QrPg  orderModel.Querier   `name:"order_pg_qr"`
	QrEs  orderModel.Querier   `name:"order_es_qr"`

	TXer  txer.TXer
	Now   func() time.Time
	NewID func() uuid.UUID
}

func New(params Params) orderModel.Service {
	return &service{
		cmdPg: params.CmdPg,
		cmdEs: params.CmdEs,
		qrPg:  params.QrPg,
		qrEs:  params.QrEs,
		tXer:  params.TXer,
		now:   params.Now,
		newID: params.NewID,
	}
}

func (s *service) Create(ctx context.Context, userID string, form *orderModel.Form) (*orderModel.Order, error) {
	item := orderModel.New(userID, s.now, s.newID)
	item.FillForm(form)

	if errTx := s.tXer.WithTX(ctx, func(ctx context.Context) error {
		if err := s.cmdPg.Create(ctx, item); err != nil {
			return fmt.Errorf("order create: %w", err)
		}

		if err := s.cmdEs.Create(ctx, item); err != nil {
			return fmt.Errorf("order create: %w", err)
		}

		return nil
	}); errTx != nil {
		return nil, errTx
	}

	return item, nil
}

func (s *service) Update(ctx context.Context, userID, id string, form *orderModel.Form) (*orderModel.Order, error) {
	item, err := s.qrPg.GetItem(ctx, option.New(orderModel.Filter{
		Status: option.New(int(orderModel.StatusCreated)),
		IDs:    option.New([]string{id}),
	}))
	if err != nil {
		return nil, fmt.Errorf("get item: %w", err)
	}

	if item.UserID != userID {
		return nil, model.ErrPermissionDenied
	}

	item.FillForm(form)
	item.TSModify = s.now()

	if errTx := s.tXer.WithTX(ctx, func(ctx context.Context) error {
		if err = s.cmdPg.Update(ctx, item); err != nil {
			return fmt.Errorf("order update: %w", err)
		}

		if err = s.cmdEs.Update(ctx, item); err != nil {
			return fmt.Errorf("order update: %w", err)
		}

		return nil
	}); errTx != nil {
		return nil, errTx
	}

	return item, nil
}

func (s *service) SoftDelete(ctx context.Context, userID, id string) error {
	item, err := s.qrPg.GetItem(ctx, option.New(orderModel.Filter{
		Status: option.New(int(orderModel.StatusCreated)),
		IDs:    option.New([]string{id}),
	}))
	if err != nil {
		return fmt.Errorf("get item: %w", err)
	}

	if item.UserID != userID {
		return model.ErrPermissionDenied
	}

	item.Status = orderModel.StatusDeleted
	item.TSModify = s.now()

	if errTx := s.tXer.WithTX(ctx, func(ctx context.Context) error {
		if err = s.cmdPg.Update(ctx, item); err != nil {
			return fmt.Errorf("order update: %w", err)
		}

		if err = s.cmdEs.Update(ctx, item); err != nil {
			return fmt.Errorf("order update: %w", err)
		}

		return nil
	}); errTx != nil {
		return errTx
	}

	return nil
}

func (s *service) GetList(ctx context.Context, userID string, req *orderModel.GetListRequest) ([]*orderModel.Order, error) {
	filter := s.prepareListCondition(userID, req)

	if filter.Q.IsSet() {
		return s.qrEs.GetList(ctx, option.New(*filter), req.Orders, req.Pagination)
	}

	return s.qrPg.GetList(ctx, option.New(*filter), req.Orders, req.Pagination)
}

func (s *service) GetItem(ctx context.Context, userID string, id string) (*orderModel.Order, error) {
	item, err := s.qrPg.GetItem(ctx, option.New(orderModel.Filter{
		Status: option.New(int(orderModel.StatusCreated)),
		IDs:    option.New([]string{id}),
	}))
	if err != nil {
		return nil, fmt.Errorf("get item: %w", err)
	}

	if item.UserID != userID {
		return nil, model.ErrPermissionDenied
	}

	return item, nil
}

func (s *service) Count(ctx context.Context, userID string, req *orderModel.GetCountRequest) (int, error) {
	filter := s.prepareCountCondition(userID, req)

	if filter.Q.IsSet() {
		return s.qrEs.Count(ctx, option.New(*filter))
	}

	return s.qrPg.Count(ctx, option.New(*filter))
}

func (s *service) InnerGetItem(ctx context.Context, req *orderModel.InnerGetItemRequest) (*orderModel.Order, error) {
	item, err := s.qrPg.GetItem(ctx, option.New(orderModel.Filter{
		IDs:    req.IDs,
		UserID: req.UserID,
	}))
	if err != nil {
		return nil, fmt.Errorf("get item: %w", err)
	}

	return item, nil
}

func (s *service) InnerGetList(ctx context.Context, req *orderModel.InnerGetListRequest) ([]*orderModel.Order, error) {
	filter := s.prepareInnerListCondition(req)

	return s.qrPg.GetList(ctx, option.New(*filter), req.Orders, req.Pagination)
}

func (s *service) prepareInnerListCondition(req *orderModel.InnerGetListRequest) *orderModel.Filter {
	if req == nil {
		return &orderModel.Filter{}
	}

	return &orderModel.Filter{
		UserID: req.UserID,
		IDs:    req.IDs,
	}
}

func (s *service) prepareListCondition(userID string, req *orderModel.GetListRequest) *orderModel.Filter {
	filter := &orderModel.Filter{
		Status: option.New(int(orderModel.StatusCreated)),
		UserID: option.New(userID),
	}

	if req == nil {
		return filter
	}

	filter.Q = req.Q
	filter.IDs = req.IDs

	return filter
}

func (s *service) prepareCountCondition(userID string, req *orderModel.GetCountRequest) *orderModel.Filter {
	filter := &orderModel.Filter{
		Status: option.New(int(orderModel.StatusCreated)),
		UserID: option.New(userID),
	}

	if req == nil {
		return filter
	}

	filter.Q = req.Q
	filter.IDs = req.IDs

	return filter
}
