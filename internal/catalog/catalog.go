package catalog

import "errors"

// Product is represent object for product catalog
type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int64
}

func (c *Product) Validate() error {
	if c.Name == "" {
		return errors.New("catalog name cannot be emtpy")
	}
	return nil
}
