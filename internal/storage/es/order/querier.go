package order

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/pkg/clients/es"
	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
	"github.com/olivere/elastic/v7"
)

type querier struct {
	esCli es.Client
}

func NewQuerier(esCli es.Client) orderModel.Querier {
	return &querier{
		esCli: esCli,
	}
}

func (q *querier) GetItem(ctx context.Context, filter option.Option[orderModel.Filter]) (*orderModel.Order, error) {
	boolQuery := q.prepareQuery(filter)

	res, err := q.esCli.GetSearch(ctx, &es.GetSearchRequest{
		Index:         indexName,
		Query:         boolQuery,
		IncludeFields: option.New([]string{"id", "status", "name", "description"}),
	})
	if err != nil {
		return nil, err
	}

	if res == nil || res.Total == 0 {
		return nil, model.ErrNotFound
	}

	if res.Total > 1 {
		return nil, model.ErrMultiItems
	}

	orderDto := dto{}

	if err = json.Unmarshal(res.Result[0].Source, &orderDto); err != nil {
		return nil, err
	}

	return orderDto.toModel(), nil
}

func (q *querier) Count(ctx context.Context, filter option.Option[orderModel.Filter]) (int, error) {
	boolQuery := q.prepareQuery(filter)

	res, err := q.esCli.GetCount(ctx, &es.GetCountRequest{
		Index: indexName,
		Query: boolQuery,
	})
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (q *querier) GetList(ctx context.Context, filter option.Option[orderModel.Filter], orders option.Option[[]*order.Order], pagination option.Option[paginator.Pagination]) ([]*orderModel.Order, error) {
	objects := make([]*orderModel.Order, 0)

	boolQuery := q.prepareQuery(filter)

	esOrders, err := q.prepareOrder(orders.Value())

	res, err := q.esCli.GetSearch(ctx, &es.GetSearchRequest{
		Index:         indexName,
		Query:         boolQuery,
		IncludeFields: option.New([]string{"id", "status", "name", "description"}),
		Orders:        option.New(esOrders),
		Pagination:    pagination,
	})
	if err != nil {
		return nil, err
	}

	for _, val := range res.Result {
		orderDto := dto{}

		if err = json.Unmarshal(val.Source, &orderDto); err != nil {
			return nil, err
		}

		objects = append(objects, orderDto.toModel())
	}

	return objects, nil
}

func (q *querier) prepareQuery(filter option.Option[orderModel.Filter]) *elastic.BoolQuery {
	boolQuery := elastic.NewBoolQuery()

	if !filter.IsSet() {
		return boolQuery
	}

	filterValue := filter.Value()
	subQueries := make([]elastic.Query, 0)

	if len(filterValue.IDs.Value()) > 0 {
		ids := make([]interface{}, 0)

		for _, id := range filterValue.IDs.Value() {
			ids = append(ids, id)
		}

		subQueries = append(subQueries, elastic.NewTermsQuery("id", ids...))
	}

	if filterValue.Status.IsSet() {
		subQueries = append(subQueries, elastic.NewTermQuery("status", filterValue.Status.Value()))
	}

	if filterValue.UserID.IsSet() {
		subQueries = append(subQueries, elastic.NewTermQuery("user_id", filterValue.UserID.Value()))
	}

	if filterValue.Q.IsSet() {
		value := filterValue.Q.Value()

		prefixNameQuery := elastic.NewPrefixQuery("name.keyword", value).
			Boost(18).
			CaseInsensitive(true)

		matchNameQuery := elastic.NewMatchQuery("name", value).
			Boost(3).
			MaxExpansions(10).
			PrefixLength(5).
			Fuzziness("AUTO")

		matchDescriptionQuery := elastic.NewMatchQuery("description", value).
			MaxExpansions(50).
			PrefixLength(5).
			Fuzziness("AUTO")

		searchQuery := elastic.NewBoolQuery().Should(
			prefixNameQuery,
			matchNameQuery,
			matchDescriptionQuery)

		subQueries = append(subQueries, searchQuery)
	}

	if len(subQueries) > 0 {
		boolQuery.Must(subQueries...)
	}

	return boolQuery
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
