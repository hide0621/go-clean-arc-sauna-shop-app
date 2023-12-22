package products

import (
	"go-clean-arc-sauna-shop-app/app/application/product"
)

type handler struct {
	saveProductUseCase  *product.SaveProductUseCase
	fetchProductUseCase *product.FetchProductUseCase
}

func NewHandler(
	saveProductUseCase *product.SaveProductUseCase,
	fetchProductUseCase *product.FetchProductUseCase,
) handler {
	return handler{
		saveProductUseCase:  saveProductUseCase,
		fetchProductUseCase: fetchProductUseCase,
	}
}

type PostProductsParams struct {
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
}
