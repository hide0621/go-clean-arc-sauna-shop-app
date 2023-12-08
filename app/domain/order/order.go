package order

import (
	"context"
	cartDomain "go-clean-arc-sauna-shop-app/app/domain/cart"
	errDomain "go-clean-arc-sauna-shop-app/app/domain/error"
	"go-clean-arc-sauna-shop-app/pkg/ulid"
	"time"
)

type Order struct {
	id          string
	userID      string
	totalAmount int64
	products    OrderProducts
	orderedAt   time.Time
}

type OrderProducts []OrderProduct

type OrderProduct struct {
	productID string
	price     int64
	quantity  int
}

// ユースケースのテスト容易性のためインターフェース化する(モックオブジェクトを作れる)
type OrderDomainService interface {
	OrderProducts(ctx context.Context, cart *cartDomain.Cart, now time.Time) (string, error)
}

func NewOrderProduct(productID string, price int64, quantity int) (*OrderProduct, error) {
	// 商品IDのバリデーション
	if !ulid.IsValid(productID) {
		return nil, errDomain.NewError("商品IDの値が不正です。")
	}

	// 購入数のバリデーション
	if quantity < 1 {
		return nil, errDomain.NewError("購入数の値が不正です。")
	}

	return &OrderProduct{
		productID: productID,
		price:     price,
		quantity:  quantity,
	}, nil
}
