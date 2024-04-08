package order

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/pkg/clients/database"
	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
)

type querier struct {
	tXer *database.TXer
}

func NewQuerier(tXer *database.TXer) orderModel.Querier {
	return &querier{
		tXer: tXer,
	}
}

func (q *querier) GetItem(ctx context.Context, filter option.Option[orderModel.Filter]) (*orderModel.Order, error) {
	sb := pgBuilder.Select(newDto().columns()...)
	sb = q.prepareBase(sb, filter)

	sql, args, errPrep := sb.ToSql()
	if errPrep != nil {
		return nil, fmt.Errorf("prepare query: %w", errPrep)
	}

	d := newDto()

	if err := q.tXer.BeginFunc(ctx, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("query: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			if d.id != "" {
				return model.ErrMultiItems
			}

			if err = rows.Scan(d.values()...); err != nil {
				return fmt.Errorf("scan: %w", err)
			}
		}

		if d.id == "" {
			return model.ErrNotFound
		}

		return rows.Err()
	}); err != nil {
		return nil, err
	}

	return d.toModel(), nil
}

func (q *querier) GetList(ctx context.Context, filter option.Option[orderModel.Filter], orders option.Option[[]*order.Order], pagination option.Option[paginator.Pagination]) ([]*orderModel.Order, error) {
	sb := pgBuilder.Select(newDto().columns()...)
	sb, errPrep := q.prepare(sb, filter, orders, pagination)
	if errPrep != nil {
		return nil, fmt.Errorf("prepare query: %w", errPrep)
	}

	sql, args, errPrep := sb.ToSql()
	if errPrep != nil {
		return nil, fmt.Errorf("prepare query: %w", errPrep)
	}

	var res []*orderModel.Order

	if err := q.tXer.BeginFunc(ctx, func(tx pgx.Tx) error {
		rows, err := tx.Query(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("query: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			d := newDto()

			if err = rows.Scan(d.values()...); err != nil {
				return fmt.Errorf("scan: %w", err)
			}

			res = append(res, d.toModel())
		}

		return rows.Err()
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func (q *querier) Count(ctx context.Context, filter option.Option[orderModel.Filter]) (int, error) {
	sb := pgBuilder.Select("COUNT(*)")
	sb = q.prepareBase(sb, filter)

	sql, args, errPrep := sb.ToSql()
	if errPrep != nil {
		return 0, fmt.Errorf("prepare query: %w", errPrep)
	}

	var count int

	if err := q.tXer.BeginFunc(ctx, func(tx pgx.Tx) error {
		row := tx.QueryRow(ctx, sql, args...)

		if err := row.Scan(&count); err != nil {
			return fmt.Errorf("scan: %w", err)
		}

		return nil
	}); err != nil {
		return 0, err
	}

	return count, nil
}

var pgBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

func (q *querier) prepare(sb squirrel.SelectBuilder, filter option.Option[orderModel.Filter], orders option.Option[[]*order.Order], pagination option.Option[paginator.Pagination]) (squirrel.SelectBuilder, error) {
	sb = q.prepareBase(sb, filter)

	if !filter.IsSet() {
		return sb, nil
	}

	if orders.IsSet() {
		var ordersStr []string

		for _, val := range orders.Value() {
			ordersStr = append(ordersStr, val.Column+" "+val.Direction)
		}

		sb = sb.OrderBy(ordersStr...)
	}

	if pagination.IsSet() {
		sb = sb.Offset(uint64(pagination.Value().Offset)).Limit(uint64(pagination.Value().Limit))
	}

	return sb, nil
}

func (q *querier) prepareBase(builder squirrel.SelectBuilder, filter option.Option[orderModel.Filter]) squirrel.SelectBuilder {
	where := squirrel.And{}

	if filter.IsSet() {
		filterValue := filter.Value()
		if filterValue.IDs.IsSet() {
			where = append(where, squirrel.Eq{"id": filterValue.IDs.Value()})
		}

		if filterValue.Status.IsSet() {
			where = append(where, squirrel.Eq{"status": filterValue.Status.Value()})
		}

		if filterValue.UserID.IsSet() {
			where = append(where, squirrel.Eq{"user_id": filterValue.UserID.Value()})
		}
	}

	builder = builder.From(tableName)
	builder = builder.Where(where)

	return builder
}
