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

func NewOrder(order OrderRepository, payment PaymentRepository) *Order {
	return &Order{
		order:   order,
		payment: payment,
	}
}

func orderToRegisterPayment(orders order.Order) order.Payment {
	var itemDetails = make([]order.ItemDetail, 0)
	for _, item := range orders.Items {
		itemDetails = append(itemDetails, order.ItemDetail{
			Name:     item.Name,
			Quantity: float64(item.Qty),
			Price:    item.BasePrice,
		})
	}

	return order.Payment{
		TransactionDetail: order.TransactionDetail{
			TrxID:       orders.TrxID,
			FinalAmount: 0.0,
		},
		CustomerDetail: order.CustomerDetail{
			Name:    orders.Billing.Name,
			Address: orders.Billing.Name,
		},
		ItemDetails: itemDetails,
	}
}

// Checkout is lock items to order before doing payment transaction
func (o *Order) Checkout(ctx context.Context, orders order.Order) (order.Order, error) {
	orders.GenerateID()
	orders.SetStatus(order.STATUS_DRAFT)

	// 1. Register to payment transaction
	paymentTrxID, err := o.payment.Register(ctx, orderToRegisterPayment(orders))
	if err != nil {
		return orders, err
	}
	orders.SetPaymentTrxID(paymentTrxID)

	// 2. Store locking items / register payment
	// TBD: should we separating to another table?
	_, err = o.order.Create(ctx, orders)
	if err != nil {
		return orders, err
	}

	// 3. Redirect to payment service page
	return orders, nil
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
