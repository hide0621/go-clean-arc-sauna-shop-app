package order

import (
	cartDomain "go-clean-arc-sauna-shop-app/app/domain/cart"
	orderDomain "go-clean-arc-sauna-shop-app/app/domain/order"
)

type SaveOrderUseCase struct {
	orderDomainService orderDomain.OrderDomainService
	cartRepo           cartDomain.CartRepository
}

func NewSaveOrderUseCase(
	orderDomainService orderDomain.OrderDomainService,
	cartRepo cartDomain.CartRepository,
) *SaveOrderUseCase {

	return &SaveOrderUseCase{
		orderDomainService: orderDomainService,
		cartRepo:           cartRepo,
	}
}
