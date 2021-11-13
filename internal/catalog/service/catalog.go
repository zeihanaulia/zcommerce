package service

import (
	"context"

	"github.com/zeihanaulia/zcommerce/internal/catalog"
)

type CatalogRepository interface {
	Upsert(ctx context.Context, params catalog.CreateParams) (catalog.Product, error)
	Delete(ctx context.Context, id int64) error
}

type Catalog struct {
	repo CatalogRepository
}

func NewOrder(repo CatalogRepository) *Catalog {
	return &Catalog{repo: repo}
}

func (c *Catalog) Upsert(ctx context.Context, params catalog.CreateParams) (catalog.Product, error) {
	if err := params.Validate(); err != nil {
		return catalog.Product{}, err
	}

	product, err := c.repo.Upsert(ctx, params)
	if err != nil {
		return catalog.Product{}, err
	}

	return product, nil
}

func (c *Catalog) Delete(ctx context.Context, id int64) error {
	if err := c.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
