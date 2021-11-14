package service

import (
	"context"

	"github.com/zeihanaulia/zcommerce/internal/payment"
)

type PaymentRepository interface {
	Register(ctx context.Context, params payment.Payment) error
}

type Payment struct {
	payment PaymentRepository
}

func NewPayment(payment PaymentRepository) *Payment {
	return &Payment{payment: payment}
}

func (p *Payment) Register(ctx context.Context, payments payment.Payment) (string, error) {
	payments.GenerateTrxID()
	payments.SumFinalAmount()
	if err := p.payment.Register(ctx, payments); err != nil {
		return "", err
	}
	return payments.TransactionDetail.PaymentTrxID, nil
}
