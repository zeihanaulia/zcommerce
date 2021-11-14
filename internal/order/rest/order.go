package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zeihanaulia/zcommerce/internal/order"
)

type OrderService interface {
	Placed(ctx context.Context, paymentTrxID string) (order.Order, error)
	Checkout(ctx context.Context, args order.Order) (order.Order, error)
}

type OrderHandler struct {
	svc OrderService
}

func NewOrderHandler(svc OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (o *OrderHandler) Register(r chi.Router) {
	r.Route("/order", func(r chi.Router) {
		r.Post("/checkout", o.checkout)
		r.Post("/placed", o.placed)
	})
}

func toItemDetail(req CreateOrderRequest) []order.Item {
	var itemDetails = make([]order.Item, 0)
	for _, i := range req.Items {
		itemDetails = append(itemDetails, order.Item{
			ID:        i.ID,
			Name:      i.Name,
			Qty:       i.Qty,
			Uom:       i.Uom,
			BasePrice: i.Price,
		})
	}
	return itemDetails
}

// CreateOrderRequest defines request for creating order
type CreateOrderRequest struct {
	Items []struct {
		ID     string  `json:"id"`
		ShopID int64   `json:"shop_id"`
		SKU    string  `json:"sku"`
		Name   string  `json:"name"`
		Uom    string  `json:"uom"`
		Qty    int64   `json:"qty"`
		Price  float64 `json:"price"`
	}
	Billing struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}
}

func (c *CreateOrderRequest) GetItems() []order.Item {
	items := make([]order.Item, len(c.Items))
	for _, ii := range c.Items {
		items = append(items, order.Item{
			ID:        ii.ID,
			ShopID:    ii.ShopID,
			SKU:       ii.SKU,
			Name:      ii.Name,
			Uom:       ii.Uom,
			Qty:       ii.Qty,
			BasePrice: ii.Price,
		})
	}
	return items
}

func (c *CreateOrderRequest) GetBilling() order.Billing {
	return order.Billing{
		ID:      c.Billing.ID,
		Name:    c.Billing.Name,
		Address: c.Billing.Address,
	}
}

type CreateOrderResponse struct {
	Data struct {
		TrxID        string `json:"trx_id"`
		PaymentTrxID string `json:"paymeny_trx_id"`
	} `json:"data"`
}

func (o *OrderHandler) checkout(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		return
	}

	itemDetails := toItemDetail(req)
	resp, err := o.svc.Checkout(r.Context(), order.Order{
		Items: itemDetails,
		Billing: order.Billing{
			Name:    req.Billing.Name,
			Address: req.Billing.Address,
		},
		Status: "draft",
	})
	if err != nil {
		log.Println(err)
		return
	}

	redirect := fmt.Sprintf("http://localhost:8003/payment/%s", resp.PaymentTrxID)
	fmt.Println(redirect)
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}

type OrderPlacedRequest struct {
	PaymentTrxID string `json:"payment_trx_id"`
}

func (o *OrderHandler) placed(w http.ResponseWriter, r *http.Request) {
	var req OrderPlacedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("decoder.order.placed: ", err)
		return
	}
	fmt.Println("Order Placed")
	resp, err := o.svc.Placed(r.Context(), req.PaymentTrxID)
	if err != nil {
		fmt.Println("scv.order.placed: ", err)
		return
	}

	renderResponse(w, CreateOrderResponse{
		Data: struct {
			TrxID        string "json:\"trx_id\""
			PaymentTrxID string "json:\"paymeny_trx_id\""
		}{
			TrxID:        resp.TrxID,
			PaymentTrxID: resp.PaymentTrxID,
		},
	}, http.StatusOK)
}
