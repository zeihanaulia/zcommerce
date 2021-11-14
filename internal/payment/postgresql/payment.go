package postgresql

import (
	"context"
	"encoding/json"

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

func (p *Payment) FindByPaymentTrxID(ctx context.Context, paymentTrxID string) (payment.Payment, error) {
	res, err := p.q.SelectPayments(ctx, paymentTrxID)
	if err != nil {
		return payment.Payment{}, err
	}

	b, err := res.Payloads.MarshalJSON()
	if err != nil {
		return payment.Payment{}, err
	}

	var obj payment.Payment
	if err := json.Unmarshal(b, &obj); err != nil {
		return payment.Payment{}, err
	}

	return payment.Payment{
		TransactionDetail: payment.TransactionDetail{
			TrxID:        res.ReferenceTrxID,
			PaymentTrxID: res.TrxID,
			FinalAmount:  float64(res.FinalAmount.Exp),
		},
		CustomerDetail: payment.CustomerDetail{
			Name:    obj.CustomerDetail.Name,
			Address: obj.CustomerDetail.Address,
		},
		ItemDetails: obj.ItemDetails,
	}, nil
}

func (p *Payment) PaidPayment(ctx context.Context, paymentTrxID string) error {
	_, err := p.q.PaidPayment(ctx, db.PaidPaymentParams{
		Types:  "opo",
		Status: "paid",
		TrxID:  paymentTrxID,
	})
	return err
}
