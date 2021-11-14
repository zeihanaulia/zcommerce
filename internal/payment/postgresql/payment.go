package postgresql

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/zeihanaulia/zcommerce/internal/payment"
	"github.com/zeihanaulia/zcommerce/internal/payment/postgresql/db"
)

// Payment represent orders tables
type Payment struct {
	q *db.Queries
}

func NewPayment(d db.DBTX) *Payment {
	return &Payment{
		q: db.New(d),
	}
}

// Register
func (p *Payment) Register(ctx context.Context, payments payment.Payment) error {
	payloads, _ := payments.ToJSON()
	jsn := pgtype.JSON{}
	jsn.Set(payloads)

	finalAmount := pgtype.Numeric{}
	finalAmount.Set(payments.TransactionDetail.FinalAmount)

	_, err := p.q.RegisterPayments(ctx, db.RegisterPaymentsParams{
		TrxID:          payments.TransactionDetail.PaymentTrxID,
		ReferenceTrxID: payments.TransactionDetail.TrxID,
		FinalAmount:    finalAmount,
		Payload:        jsn,
		Status:         "new",
	})

	return err
}
