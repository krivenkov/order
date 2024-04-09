package order

import (
	"time"

	"github.com/krivenkov/order/internal/model/order"
)

func init() {
	d := newDto()
	if len(d.columns()) != len(d.values()) {
		panic("order.item.dto: len(columns) != len(values)")
	}
}

const (
	tableName = `"order".items`
)

type dto struct {
	id       string
	tsCreate time.Time
	tsModify time.Time
	status   int64

	userID      string
	name        string
	description string
}

func newDto() *dto {
	return &dto{}
}

func (d *dto) columns() []string {
	return []string{"id", "ts_create", "ts_modify", "status", "user_id", "name", "description"}
}

func (d *dto) values() []interface{} {
	return []interface{}{&d.id, &d.tsCreate, &d.tsModify, &d.status, &d.userID, &d.name, &d.description}
}

func (d *dto) toMap() map[string]interface{} {
	columns, values := d.columns(), d.values()

	dm := make(map[string]interface{}, len(columns))
	for i, c := range columns {
		dm[c] = values[i]
	}
	return dm
}

func (d *dto) toModel() *order.Order {
	return &order.Order{
		ID:          d.id,
		TSCreate:    d.tsCreate,
		TSModify:    d.tsModify,
		Status:      order.Status(d.status),
		UserID:      d.userID,
		Name:        d.name,
		Description: d.description,
	}
}

func (d *dto) fromModel(source *order.Order) {
	target := dto{
		id:          source.ID,
		tsCreate:    source.TSCreate,
		tsModify:    source.TSModify,
		status:      int64(source.Status),
		userID:      source.UserID,
		name:        source.Name,
		description: source.Description,
	}

	*d = target
}
