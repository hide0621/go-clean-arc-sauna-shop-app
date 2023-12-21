package repository

import (
	"context"
	"go-clean-arc-sauna-shop-app/app/domain/product"
	"go-clean-arc-sauna-shop-app/app/infrastructure/mysql/db"
	"go-clean-arc-sauna-shop-app/app/infrastructure/mysql/db/dbgen"
)

type productRepository struct{}

func NewProductRepository() *product.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Save(ctx context.Context, product *product.Product) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertProduct(ctx, dbgen.UpsertProductParams{
		ID:          product.ID(),
		OwnerID:     product.OwnerID(),
		Name:        product.Name(),
		Description: product.Description(),
		Price:       product.Price(),
		Stock:       int32(product.Stock()),
	}); err != nil {
		return err
	}
	return nil
}
