package order

import (
	"context"

	"github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/pkg/clients/es"
)

type commander struct {
	esCli es.Client
}

func NewCommander(esCli es.Client) order.Commander {
	return &commander{
		esCli: esCli,
	}
}

func (c *commander) Create(ctx context.Context, item *order.Order) error {
	d := newDto()
	d.fromModel(item)

	return c.esCli.Save(ctx, &es.SaveRequest{
		Index:   indexName,
		ID:      item.ID,
		Obj:     d,
		Refresh: es.RefreshTypeWaitFor,
	})
}

func (c *commander) Update(ctx context.Context, item *order.Order) error {
	d := newDto()
	d.fromModel(item)

	return c.esCli.Save(ctx, &es.SaveRequest{
		Index:   indexName,
		ID:      item.ID,
		Obj:     d,
		Refresh: es.RefreshTypeWaitFor,
	})
}

func (c *commander) Delete(ctx context.Context, item *order.Order) error {
	return c.esCli.DeleteByID(ctx, indexName, item.ID)
}
