package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zeihanaulia/zcommerce/internal/payment"
)

type PaymentService interface {
	Register(ctx context.Context, params payment.Payment) (string, error)
}

type PaymentHandler struct {
	payment PaymentService
}

func NewPaymentHandler(payment PaymentService) *PaymentHandler {
	return &PaymentHandler{payment: payment}
}

func (p *PaymentHandler) Register(r chi.Router) {
	r.Route("/payment", func(r chi.Router) {
		r.Get("/{trxId}", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Println("payment page")
		})
		r.Post("/register", p.register)
	})
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
	ItemDetails []struct {
		ID       int64   `json:"id"`
		Name     string  `json:"name"`
		Quantity int64   `json:"quantity"`
		Price    float64 `json:"price"`
	} `json:"item_details"`
}

func toItemDetails(req RegisterRequest) []payment.ItemDetail {
	var itemDetails = make([]payment.ItemDetail, 0)
	for _, item := range req.ItemDetails {
		itemDetails = append(itemDetails, payment.ItemDetail{
			ID:       "",
			Name:     item.Name,
			Quantity: float64(item.Quantity),
			Price:    item.Price,
		})
	}
	return itemDetails
}

type RegisterResponse struct {
	Data struct {
		PaymentTrxID string `json:"payment_trx_id"`
	} `json:"data"`
}

func (p *PaymentHandler) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		return
	}
	itemDetails := toItemDetails(req)
	paymentTrxID, err := p.payment.Register(r.Context(), payment.Payment{
		TransactionDetail: payment.TransactionDetail{
			TrxID:       req.TransactionDetail.TrxID,
			FinalAmount: req.TransactionDetail.FinalAmount,
		},
		CustomerDetail: payment.CustomerDetail{
			Name:    req.CustomerDetail.Name,
			Address: req.CustomerDetail.Address,
		},
		ItemDetails: itemDetails,
	})
	if err != nil {
		return
	}

	renderResponse(w, RegisterResponse{
		Data: struct {
			PaymentTrxID string "json:\"payment_trx_id\""
		}{
			PaymentTrxID: paymentTrxID,
		},
	}, http.StatusOK)
}
