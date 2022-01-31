package service

import (
	"context"
	"fmt"

	"github.com/zeihanaulia/zcommerce/internal/payment"
	"go.elastic.co/apm"
)

type OPORepository interface {
	Paid(ctx context.Context, phoneNumber string) (string, error)
}

type PaymentRepository interface {
	Register(ctx context.Context, params payment.Payment) error
	FindByPaymentTrxID(ctx context.Context, paymentTrxID string) (payment.Payment, error)
	PaidPayment(ctx context.Context, paymentTrxID string) error
}

type OrderRepository interface {
	Placed(ctx context.Context, paymentTrxID string) error
}

type Payment struct {
	payment PaymentRepository
	opo     OPORepository
	order   OrderRepository
}

func NewPayment(payment PaymentRepository, opo OPORepository, order OrderRepository) *Payment {
	return &Payment{payment: payment, opo: opo, order: order}
}

func (p *Payment) Register(ctx context.Context, payments payment.Payment) (string, error) {
	span, ctx := apm.StartSpan(ctx, "Payment.Register", "custom")
	defer span.End()

	payments.GenerateTrxID()
	payments.SumFinalAmount()
	if err := p.payment.Register(ctx, payments); err != nil {
		return "", err
	}
	return payments.TransactionDetail.PaymentTrxID, nil
}

func (p *Payment) ByPaymentTrxID(ctx context.Context, paymentTrxID string) (payment.Payment, error) {
	span, ctx := apm.StartSpan(ctx, "Payment.ByPaymentTrxID", "custom")
	defer span.End()

	return p.payment.FindByPaymentTrxID(ctx, paymentTrxID)
}

func (p *Payment) OPOPaid(ctx context.Context, paymentTrxID string) error {
	span, ctx := apm.StartSpan(ctx, "Payment.OPOPaid", "custom")
	defer span.End()

	// 1. Send to opo
	opoTrxID, err := p.opo.Paid(ctx, paymentTrxID)
	if err != nil {
		return err
	}

	// 2. Update status
	if err := p.payment.PaidPayment(ctx, paymentTrxID); err != nil {
		return err
	}

	fmt.Println(opoTrxID)

	// 2. Callback to order service
	if err := p.order.Placed(ctx, paymentTrxID); err != nil {
		return err
	}

	return nil
}
