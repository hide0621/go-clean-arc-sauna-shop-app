package order

import (
	"context"
	cartDomain "go-clean-arc-sauna-shop-app/app/domain/cart"
	errDomain "go-clean-arc-sauna-shop-app/app/domain/error"
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

func (ds *orderDomainService) OrderProduct(
	ctx context.Context,
	cart *cartDomain.Cart,
	now time.Time,
) (string, error) {

	// 注文する商品情報の取得
	ps, err := ds.productRepo.FindByIDs(ctx, cart.ProductIDs())
	if err != nil {
		return "", nil
	}

	productMap := make(map[string]*productDomain.Product)
	for _, p := range ps {
		productMap[p.ID()] = p
	}

	// 注文を行う
	ops := make([]OrderProduct, 0, len(cart.ProductIDs()))
	for _, cp := range cart.Products() {
		p, ok := productMap[cp.ProductID()]
		op, err := NewOrderProduct(cp.ProductID(), p.Price(), cp.Quantity())
		if err != nil {
			return "", err
		}
		ops = append(ops, *op)
		if !ok {
			// 注文した商品の商品詳細が見つからない場合はエラー
			// 商品を注文すると同時に 、商品が削除された場合等に発生
			return "", errDomain.NewError("商品が見つかりません 。")
		}
		if err := p.Consume(cp.Quantity()); err != nil {
			return "", err
		}
		if err := ds.productRepo.Save(ctx, p); err != nil {
			return "", err
		}
	}

	// 注文した商品を履歴として保存する
	o, err := NewOrder(cart.UserID(), OrderProducts(ops).TotalAmount(), ops, now)
	if err != nil {
		return "", err
	}
	if err := ds.orderRepo.Save(ctx, o); err != nil {
		return "", nil
	}

	return o.ID(), nil
}
