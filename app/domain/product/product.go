package product

import (
	errDomain "go-clean-arc-sauna-shop-app/app/domain/error" //パッケージをエイリアス化
	"go-clean-arc-sauna-shop-app/pkg/ulid"
	"unicode/utf8"
)

type Product struct {
	id          string
	ownerID     string
	name        string
	description string
	price       int64
	stock       int
}

const (
	nameLengthMin = 1
	nameLengthMax = 100

	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)

// データベースから取得した既存のIDを用いて商品オブジェクトを再構築
func Reconstruct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
		id,
		ownerID,
		name,
		description,
		price,
		stock,
	)
}

// 新しいIDを生成し、それをもとに商品オブジェクトを生成
func NewProduct(
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
		ulid.NewULID(),
		ownerID,
		name,
		description,
		price,
		stock,
	)
}

func newProduct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int) (*Product, error) {

	if !ulid.IsValid(ownerID) {
		return nil, errDomain.NewError("オーナーIDの値が不正です 。")
	}

	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("商品名の値が不正です 。")
	}

	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError("商品説明の値が不正です 。")
	}

	if price < 1 {
		return nil, errDomain.NewError("価格の値が不正です 。")
	}

	// // 在庫はないけど、商品は登録したい等あるため、0は許容する
	if stock < 0 {
		return nil, errDomain.NewError("在庫数の値が不正です 。")
	}

	return &Product{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
	}, nil
}

func (p *Product) Consume(quantity int) error {

	// こっちでは？
	if p.stock < 0 {
		errDomain.NewError("在庫数の値が不正です。")
	}

	// こっちではなく
	/*
		if quantity < 0 {
			return errDomain.NewError("在庫数の値が不正です。")
		}
	*/

	// この処理も必要では？
	if quantity < 0 {
		errDomain.NewError("注文数が不正です")
	}

	if p.stock-quantity < 0 {
		errDomain.NewError("在庫数が不足しています。")
	}

	p.stock -= quantity
	return nil
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) OwnerID() string {
	return p.ownerID
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Price() int64 {
	return p.price
}

func (p *Product) Stock() int {
	return p.stock
}
