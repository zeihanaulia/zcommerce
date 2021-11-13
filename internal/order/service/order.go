package service

import (
	"context"

	"github.com/zeihanaulia/zcommerce/internal/order"
)

type Order struct {
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) Create(ctx context.Context, paymentTrxID string, items []order.Item, billing order.Billing) (order.Order, error) {
	// Create order to db
	// Dispatch order to vendor
	return order.Order{
		ID:           "3158c537-081a-44ae-9477-011d990bc7e8",
		TrxID:        "SO1110-000001",
		PaymentTrxID: paymentTrxID,
		Items:        items,
		Billing:      billing,
	}, nil
}
