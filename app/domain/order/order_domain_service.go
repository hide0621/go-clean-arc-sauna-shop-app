package order

import (
	"context"
	cartDomain "go-clean-arc-sauna-shop-app/app/domain/cart"
	productDomain "go-clean-arc-sauna-shop-app/app/domain/product"
	"time"
)

type orderDomainService struct {
	orderRepo   OrderRepository
	productRepo productDomain.ProductRepository
}

func NewOrderDomainService(
	orderRepo OrderRepository,
	productRepo productDomain.ProductRepository,
) OrderDomainService {
	return &orderDomainService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (ds *orderDomainService) OrderProduct(ctx context.Context,
	cart *cartDomain.Cart,
	now time.Time,
) (string, error) {

	ps, err := ds.productRepo.FindByIDs(ctx, cart.ProductIDs())
	if err != nil {
		return "", nil
	}

	productMap := make(map[string]*productDomain.Product)
	for _, p := range ps {
		productMap[p.ID()] = p
	}

	// ops := make([]OrderProduct, 0, len(cart.ProductIDs()))
}
