//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
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

func NewOrder(userID string, totalAmount int64, products []OrderProduct, now time.Time) (*Order, error) {
	return newOrder(
		ulid.NewULID(),
		userID,
		totalAmount,
		products,
		now,
	)
}

func Reconstruct(id string, userID string, totalAmount int64, products []OrderProduct, OrderedAt time.Time) (*Order, error) {
	return newOrder(
		id,
		userID,
		totalAmount,
		products,
		OrderedAt,
	)
}

func newOrder(
	id string,
	userID string,
	totalAmount int64,
	products []OrderProduct,
	orderedAt time.Time,
) (*Order, error) {
	// userIDのバリデーション
	if !ulid.IsValid(userID) {
		return nil, errDomain.NewError("ユーザーIDの値が不正です。")
	}

	// 購入金額のバリデーション
	// 割引等で合計金額が0円になることはあるため、0円以上を許容とする
	if totalAmount < 0 {
		return nil, errDomain.NewError("購入金額の値が不正です。")
	}

	// 購入商品のバリデーション
	if len(products) < 1 {
		return nil, errDomain.NewError("購入商品がありません。")
	}
	return &Order{
		id:          id,
		userID:      userID,
		totalAmount: totalAmount,
		products:    products,
		orderedAt:   orderedAt,
	}, nil
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

// 合計金額計算
func (p OrderProducts) TotalAmount() int64 {
	var totalAmount int64
	for _, product := range p {
		totalAmount += product.price * int64(product.quantity)
	}
	return totalAmount
}

func (p *Order) ID() string {
	return p.id
}
