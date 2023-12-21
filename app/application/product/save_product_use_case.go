package product

import (
	productDomain "go-clean-arc-sauna-shop-app/app/domain/product"
)

type SaveProductUseCase struct {
	productRepo productDomain.ProductRepository
}

func NewSaveProductUseCase(
	productRepo productDomain.ProductRepository,
) *SaveProductUseCase {
	return &SaveProductUseCase{
		productRepo: productRepo,
	}
}
