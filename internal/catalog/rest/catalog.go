package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zeihanaulia/zcommerce/internal/catalog"
	"github.com/zeihanaulia/zcommerce/internal/order"
)

type CatalogService interface {
	Upsert(ctx context.Context, params catalog.CreateParams) (catalog.Product, error)
	Delete(ctx context.Context, id int64) error
}

type CatalogHandler struct {
	svc CatalogService
}

func NewOrderHandler(svc CatalogService) *CatalogHandler {
	return &CatalogHandler{svc: svc}
}

func (c *CatalogHandler) Register(r chi.Router) {
	r.Route("/catalog", func(r chi.Router) {
		r.Post("/upsert", c.create)
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

type UpsertCatalogResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int64   `json:"stock"`
}

func (o *CatalogHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	resp, err := o.svc.Upsert(r.Context(), catalog.CreateParams{})
	if err != nil {
		return
	}

	renderResponse(w, UpsertCatalogResponse{
		ID:    resp.ID,
		Name:  resp.Name,
		Price: resp.Price,
		Stock: resp.Stock,
	}, http.StatusOK)
}
