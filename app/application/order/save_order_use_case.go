package order

import (
	"context"
	cartDomain "go-clean-arc-sauna-shop-app/app/domain/cart"
	errDomain "go-clean-arc-sauna-shop-app/app/domain/error"
	orderDomain "go-clean-arc-sauna-shop-app/app/domain/order"
	"time"
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

type SaveOrderUseCaseInputDto struct {
	ProductID string
	Quantity  int
}

func (uc *SaveOrderUseCase) Run(
	ctx context.Context,
	userID string,
	dtos []SaveOrderUseCaseInputDto,
	now time.Time,
) (string, error) {

	cart, err := uc.getValidCart(ctx, userID, dtos)
	if err != nil {
		return "", nil
	}

	orderID, err := uc.orderDomainService.OrderProducts(ctx, cart, now)
	if err != nil {
		return "", err
	}

	return orderID, nil
}

func (uc *SaveOrderUseCase) getValidCart(
	ctx context.Context,
	userID string,
	dtos []SaveOrderUseCaseInputDto,
) (*cartDomain.Cart, error) {

	cart, err := uc.cartRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, dto := range dtos {
		pq, err := cart.QuantityByProductID(dto.ProductID)
		if err != nil {
			return nil, err
		}
		if pq != dto.Quantity {
			return nil, errDomain.NewError("カートの商品数が変更されています 。")
		}
	}

	return cart, nil
}
