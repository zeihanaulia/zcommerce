package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const (
	STATUS_DRAFT    Status = "draft"
	STATUS_PLACED   Status = "placed"
	STATUS_CANCELED Status = "canceled"
)

type Status string

// Order is represent object for creating new order
type Order struct {
	ID           string  `json:"id"`
	TrxID        string  `json:"trx_id"`
	PaymentTrxID string  `json:"payment_trx_id"`
	Items        []Item  `json:"items"`
	Billing      Billing `json:"billing"`
	Status       Status  `json:"status"`
}

func (o *Order) GenerateID() {
	now := time.Now()
	o.TrxID = fmt.Sprintf("SO-%v", now.Unix())
}

func (o *Order) SetStatus(val Status) {
	o.Status = val
}

func (o *Order) ItemsToJSON() ([]byte, error) {
	b, err := json.Marshal(o.Items)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b, nil
}

func (o *Order) Validate() error {
	if o.PaymentTrxID == "" {
		return errors.New("order payment trx id is required")
	}

	for _, item := range o.Items {
		// TBD: should we collect all error before return or not?
		if err := item.validate(); err != nil {
			return fmt.Errorf("items is invalid %w", err)
		}
	}

	if err := o.Billing.validate(); err != nil {
		return fmt.Errorf("billing is invalid %w", err)
	}

	return nil
}

// Item means the items from order
type Item struct {
	ID        string  `json:"id"`
	ShopID    int64   `json:"shop_id"`
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Uom       string  `json:"uom"`
	Qty       int64   `json:"qty"`
	BasePrice float64 `json:"base_price"`
}

func (i *Item) validate() error {
	if i.ShopID == 0 {
		return errors.New("item shop is required")
	}

	if i.SKU == "" {
		return errors.New("item sku is required")
	}

	if i.Uom == "" {
		return errors.New("item uom is required")
	}

	return nil
}

// Billing is data collection for customer billing information
type Billing struct {
	ID      string
	Name    string
	Address string
}

func (b *Billing) validate() error {
	if b.Name == "" {
		return errors.New("billing name is required")
	}

	if b.Address == "" {
		return errors.New("billing address is required")
	}

	return nil
}
