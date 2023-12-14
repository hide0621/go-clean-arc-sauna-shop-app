package repository

import (
	"go-clean-arc-sauna-shop-app/app/domain/product"
)

type productRepository struct{}

func NewProductRepository() *product.ProductRepository {
	return &productRepository{}
}
