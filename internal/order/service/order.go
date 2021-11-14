package service

import (
	"context"

	"github.com/zeihanaulia/zcommerce/internal/order"
)

type PaymentRepository interface {
	Register(ctx context.Context, args order.Payment) (string, error)
}

type OrderRepository interface {
	Create(ctx context.Context, args order.Order) (order.Order, error)
}
type Order struct {
	order   OrderRepository
	payment PaymentRepository
}

func NewOrder(order OrderRepository) *Order {
	return &Order{
		order: order,
		// payment: payment,
	}
}

// Checkout is lock items to order before doing payment transaction
func (o *Order) Checkout(ctx context.Context, orders order.Order) (string, error) {
	// 1. Register to payment transaction
	// paymentTrxID, err := o.payment.Register(ctx, order.Payment{})
	// if err != nil {
	// 	return "", err
	// }

	// 2. Store locking items / register payment
	// TBD: should we separating to another table?
	orders.GenerateID()
	orders.SetStatus(order.STATUS_DRAFT)
	_, err := o.order.Create(ctx, orders)
	if err != nil {
		return "", err
	}

	// 3. Redirect to payment service page
	return "paymentTrxID", nil
}

// Placed is creating order from payment service
func (o *Order) Placed(ctx context.Context, paymentTrxID string, items []order.Item, billing order.Billing) (order.Order, error) {
	//1. Check signature
	//2. Update status order to placed
	orders, err := o.order.Create(ctx, order.Order{})
	if err != nil {
		return order.Order{}, err
	}

	return orders, nil
}
