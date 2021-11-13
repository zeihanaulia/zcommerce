package order

import (
	"errors"
	"fmt"
)

// Order is represent object for creating new order
type Order struct {
	ID           string
	TrxID        string
	PaymentTrxID string
	Items        []Item
	Billing      Billing
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
	ID        string
	ShopID    int64
	SKU       string
	Name      string
	Uom       string
	Qty       int64
	BasePrice float64
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
