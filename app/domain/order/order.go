package order

import (
	"context"
	cartDomain "go-clean-arc-sauna-shop-app/app/domain/cart"
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
	OrderProducts(ctx context.Context, cart *cartDomain.Cart, now time.Time)
}
