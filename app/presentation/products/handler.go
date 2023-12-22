package products

import (
	"go-clean-arc-sauna-shop-app/app/application/product"
)

type handler struct {
	saveProductUseCase  *product.SaveProductUseCase
	fetchProductUseCase *product.FetchProductUseCase
}
