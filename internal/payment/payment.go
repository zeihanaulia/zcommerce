package payment

import (
	"encoding/json"
	"fmt"
	"time"
)

// Payment defines a type of payment service
// The service should to know what customer order
// and Locking items before payment
type Payment struct {
	TransactionDetail TransactionDetail `json:"transaction_detail"`
	CustomerDetail    CustomerDetail    `json:"customer_detail"`
	ItemDetails       []ItemDetail      `json:"item_details"`
}

func (t *Payment) GenerateTrxID() {
	now := time.Now()
	t.TransactionDetail.PaymentTrxID = fmt.Sprintf("PAY-%v", now.Unix())
}

func (t *Payment) SumFinalAmount() {
	var finalAmount float64
	for _, items := range t.ItemDetails {
		finalAmount += items.Quantity * items.Price
	}
	t.TransactionDetail.FinalAmount = finalAmount
}

func (p *Payment) ToJSON() ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b, nil
}

// TransactionDetail defines type of transaction detail
// such as transaction id as a key
type TransactionDetail struct {
	TrxID        string  `json:"trx_id"`
	PaymentTrxID string  `json:"payment_trx_id"`
	FinalAmount  float64 `json:"final_amount"`
}

// CustomerDetail defines type of customer detail
type CustomerDetail struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// ItemDetail defines type of items detail
type ItemDetail struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Uom      string  `json:"uom"`
	Price    float64 `json:"price"`
}
