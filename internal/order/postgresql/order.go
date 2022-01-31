package postgresql

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgtype"
	"github.com/zeihanaulia/zcommerce/internal/order"
	"github.com/zeihanaulia/zcommerce/internal/order/postgresql/db"
	"go.elastic.co/apm"
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
	span, ctx := apm.StartSpan(ctx, "Postgresql.Create", "custom")
	defer span.End()

	itmsJson, _ := params.ItemsToJSON()
	jsn := pgtype.JSON{}
	jsn.Set(itmsJson)
	id, err := o.q.CreateOrders(ctx, db.CreateOrdersParams{
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
		ID: id,
	}, nil
}

func (o *Order) UpdateStatusOrder(ctx context.Context, paymentTrxID string) error {
	span, ctx := apm.StartSpan(ctx, "Postgresql.UpdateStatusOrder", "custom")
	defer span.End()

	_, err := o.q.OrderPlaced(ctx, db.OrderPlacedParams{
		Status:       string(order.STATUS_PLACED),
		PaymentTrxID: paymentTrxID,
	})
	return err
}

func (o *Order) FindPayload(ctx context.Context, paymentTrxID string) (order.Order, error) {
	span, ctx := apm.StartSpan(ctx, "Postgresql.FindPayload", "custom")
	defer span.End()

	resp, err := o.q.SelectPayloads(ctx, paymentTrxID)
	if err != nil {
		return order.Order{}, err
	}

	b, err := resp.LockItems.MarshalJSON()
	if err != nil {
		return order.Order{}, err
	}

	var obj []order.Item
	if err := json.Unmarshal(b, &obj); err != nil {
		return order.Order{}, err
	}

	return order.Order{
		ID:    resp.ID,
		Items: obj,
	}, nil
}

func (o *Order) CreateDetail(ctx context.Context, orders order.Order) error {
	span, ctx := apm.StartSpan(ctx, "Postgresql.FindPayload", "custom")
	defer span.End()

	for _, item := range orders.Items {
		price := pgtype.Numeric{}
		price.Set(item.BasePrice)
		_, err := o.q.CreateOrdersDetail(ctx, db.CreateOrdersDetailParams{
			OrderID:  orders.ID,
			Name:     item.Name,
			Quantity: int32(item.Qty),
			Price:    price,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
