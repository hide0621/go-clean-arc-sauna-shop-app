package repository

import (
	"context"
	"go-clean-arc-sauna-shop-app/app/domain/product"
)

type productRepository struct{}

func NewProductRepository() *product.ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Save(ctx context.Context, product *product.Product) error {

}
