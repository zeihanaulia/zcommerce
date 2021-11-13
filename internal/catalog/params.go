package catalog

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// CreateParams define arguments for creting catalog product
type CreateParams struct {
	Name  string
	Price float64
	Stock int64
}

// Validate check based on prodduct entity
func (c CreateParams) Validate() error {
	p := Product{
		Name:  c.Name,
		Price: c.Price,
		Stock: c.Stock,
	}

	if err := validation.Validate(&p); err != nil {
		return err
	}

	return nil
}
