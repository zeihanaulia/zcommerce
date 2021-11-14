package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zeihanaulia/zcommerce/internal/order"
)

type Payment struct {
	host string
}

func NewPayment(host string) *Payment {
	return &Payment{host: host}
}

type RegisterRequest struct {
	TransactionDetail struct {
		TrxID       string  `json:"trx_id"`
		FinalAmount float64 `json:"final_amount"`
	} `json:"transaction_detail"`
	CustomerDetail struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"customer_detail"`
	ItemDetails []ItemDetail `json:"item_details"`
}

type ItemDetail struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
}

func paymentToRegister(payments order.Payment) RegisterRequest {
	var itemDetails = make([]ItemDetail, 0)
	for _, item := range payments.ItemDetails {
		itemDetails = append(itemDetails, ItemDetail{
			Name:     item.Name,
			Quantity: int64(item.Quantity),
			Price:    item.Price,
		})
	}

	return RegisterRequest{
		TransactionDetail: struct {
			TrxID       string  "json:\"trx_id\""
			FinalAmount float64 "json:\"final_amount\""
		}{
			TrxID:       payments.TransactionDetail.TrxID,
			FinalAmount: payments.TransactionDetail.FinalAmount,
		},
		CustomerDetail: struct {
			Name    string "json:\"name\""
			Address string "json:\"address\""
		}{
			Name:    payments.CustomerDetail.Name,
			Address: payments.CustomerDetail.Address,
		},
		ItemDetails: nil,
	}
}

type RegisterResponse struct {
	Data struct {
		PaymentTrxID string `json:"payment_trx_id"`
	} `json:"data"`
}

func (p *Payment) Register(ctx context.Context, payments order.Payment) (string, error) {
	// TODO: add tracing
	postBody, _ := json.Marshal(paymentToRegister(payments))
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(fmt.Sprintf("%s/payment/register", p.host), "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	var res RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res.Data.PaymentTrxID, nil
}
