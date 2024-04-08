package order

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/pkg/clients/database"
)

type commander struct {
	tXer *database.TXer
}

func NewCommander(tXer *database.TXer) order.Commander {
	return &commander{
		tXer: tXer,
	}
}

func (c *commander) Create(ctx context.Context, item *order.Order) error {
	ib := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert(tableName)

	d := newDto()
	d.fromModel(item)

	ib = ib.SetMap(d.toMap())

	return c.exec(ctx, ib)
}

func (c *commander) Update(ctx context.Context, item *order.Order) error {
	ub := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update(tableName)

	d := newDto()
	d.fromModel(item)
	ub = ub.SetMap(d.toMap()).
		Where(squirrel.Eq{"id": d.id})

	return c.exec(ctx, ub)
}

func (c *commander) Delete(ctx context.Context, item *order.Order) error {
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete(tableName).
		Where(squirrel.Eq{"id": item.ID})

	return c.exec(ctx, b)
}

func (c *commander) exec(ctx context.Context, sq squirrel.Sqlizer) error {
	sql, args, err := sq.ToSql()
	if err != nil {
		return fmt.Errorf("create query: %w", err)
	}

	if err = c.tXer.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, errExec := tx.Exec(ctx, sql, args...)
		return errExec
	}); err != nil {
		return err
	}

	return nil
}
