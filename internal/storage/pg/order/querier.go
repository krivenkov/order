package order

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/pkg/clients/database"
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

func (q *querier) GetItem(ctx context.Context, filter *orderModel.Filter) (*orderModel.Order, error) {
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

func (q *querier) GetList(ctx context.Context, filter *orderModel.Filter, orders []*order.Order, pagination *paginator.Pagination) ([]*orderModel.Order, error) {
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

func (q *querier) Count(ctx context.Context, filter *orderModel.Filter) (int, error) {
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

func (q *querier) prepare(sb squirrel.SelectBuilder, filter *orderModel.Filter, orders []*order.Order, pagination *paginator.Pagination) (squirrel.SelectBuilder, error) {
	sb = q.prepareBase(sb, filter)

	if len(orders) > 0 {
		prepareOrders, err := q.prepareOrder(orders)
		if err != nil {
			return sb, err
		}

		var ordersStr []string
		for _, val := range prepareOrders {
			ordersStr = append(ordersStr, val.Column+" "+val.Direction)
		}

		sb = sb.OrderBy(ordersStr...)
	}

	if pagination != nil {
		sb = sb.Offset(uint64(pagination.Offset)).Limit(uint64(pagination.Limit))
	}

	return sb, nil
}

func (q *querier) prepareBase(builder squirrel.SelectBuilder, filter *orderModel.Filter) squirrel.SelectBuilder {
	where := squirrel.And{}

	if filter != nil {
		if filter.IDs.IsSet() {
			where = append(where, squirrel.Eq{"id": filter.IDs.Value()})
		}

		if filter.Status.IsSet() {
			where = append(where, squirrel.Eq{"status": filter.Status.Value()})
		}

		if filter.UserID.IsSet() {
			where = append(where, squirrel.Eq{"user_id": filter.UserID.Value()})
		}
	}

	builder = builder.From(tableName)
	builder = builder.Where(where)

	return builder
}

func (q *querier) prepareOrder(orders []*order.Order) ([]*order.Order, error) {
	esOrders := make([]*order.Order, 0, len(orders))

	for _, val := range orders {
		switch val.Column {
		case orderModel.NameSortKey:
			esOrders = append(esOrders, &order.Order{
				Column:    nameSortKey,
				Direction: val.Direction,
			})
		case orderModel.IDSortKey:
			esOrders = append(esOrders, &order.Order{
				Column:    idSortKey,
				Direction: val.Direction,
			})
		default:
			return nil, fmt.Errorf("invalid sort column = %s", val.Column)
		}

	}

	return esOrders, nil
}
