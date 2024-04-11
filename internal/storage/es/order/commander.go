package order

import (
	"context"

	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/pkg/clients/es"
	"github.com/olivere/elastic/v7"
)

type commander struct {
	esCli es.Client
}

func NewCommander(esCli es.Client) orderModel.Commander {
	return &commander{
		esCli: esCli,
	}
}

func (c *commander) Create(ctx context.Context, item *orderModel.Order) error {
	d := newDto()
	d.fromModel(item)

	return c.esCli.Save(ctx, &es.SaveRequest{
		Index:   indexName,
		ID:      item.ID,
		Obj:     d,
		Refresh: es.RefreshTypeWaitFor,
	})
}

func (c *commander) Update(ctx context.Context, item *orderModel.Order) error {
	d := newDto()
	d.fromModel(item)

	return c.esCli.Save(ctx, &es.SaveRequest{
		Index:   indexName,
		ID:      item.ID,
		Obj:     d,
		Refresh: es.RefreshTypeWaitFor,
	})
}

func (c *commander) Delete(ctx context.Context, item *orderModel.Order) error {
	return c.esCli.DeleteByID(ctx, indexName, item.ID)
}

func (c *commander) Disable(ctx context.Context, userID string) error {
	return c.esCli.UpdateByScript(ctx, &es.UpdateByScriptRequest{
		Index:   indexName,
		Refresh: es.RefreshTypeWaitFor,
		Query:   elastic.NewTermQuery("user", userID),
		Script:  elastic.NewScriptInline("ctx._source.status = 2").Lang("painless"),
	})
}
