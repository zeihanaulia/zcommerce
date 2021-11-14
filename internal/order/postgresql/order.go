package postgresql

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/zeihanaulia/zcommerce/internal/order"
	"github.com/zeihanaulia/zcommerce/internal/order/postgresql/db"
)

// Order represent orders tables
type Order struct {
	q *db.Queries
}

func NewOrder(d db.DBTX) *Order {
	return &Order{
		q: db.New(d),
	}
}

// Create
func (o *Order) Create(ctx context.Context, params order.Order) (order.Order, error) {
	itmsJson, _ := params.ItemsToJSON()
	jsn := pgtype.JSON{}
	jsn.Set(itmsJson)
	id, err := o.q.OrdersTask(ctx, db.OrdersTaskParams{
		TrxID:           params.TrxID,
		PaymentTrxID:    params.PaymentTrxID,
		Status:          string(params.Status),
		LockItems:       jsn,
		CustomerName:    params.Billing.Name,
		CustomerAddress: params.Billing.Address,
	})
	if err != nil {
		return order.Order{}, err
	}
	return order.Order{
		ID: string(id),
	}, nil
}
