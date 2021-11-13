package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zeihanaulia/zcommerce/internal/order"
)

type OrderService interface {
	Create(ctx context.Context, paymentTrxID string, items []order.Item, billing order.Billing) (order.Order, error)
}

type OrderHandler struct {
	svc OrderService
}

func NewOrderHandler(svc OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (o *OrderHandler) Register(r chi.Router) {
	r.Route("/order", func(r chi.Router) {
		r.Post("/publish", o.create)
	})
}

// CreateOrderRequest defines request for creating order
type CreateOrderRequest struct {
	PaymentTrxID string `json:"payment_trx_id"`
	Items        []struct {
		ID        string  `json:"id"`
		ShopID    int64   `json:"shop_id"`
		SKU       string  `json:"sku"`
		Name      string  `json:"name"`
		Uom       string  `json:"uom"`
		Qty       int64   `json:"qty"`
		BasePrice float64 `json:"base_price"`
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
			BasePrice: ii.BasePrice,
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
	ID           string `json:"id"`
	TrxID        string `json:"trx_id"`
	PaymentTrxID string `json:"paymeny_trx_id"`
}

func (o *OrderHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	resp, err := o.svc.Create(r.Context(), req.PaymentTrxID, req.GetItems(), req.GetBilling())
	if err != nil {
		return
	}

	renderResponse(w, CreateOrderResponse{
		ID:           resp.ID,
		TrxID:        resp.TrxID,
		PaymentTrxID: resp.PaymentTrxID,
	}, http.StatusOK)
}
