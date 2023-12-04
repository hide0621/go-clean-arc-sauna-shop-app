package product

import (
	errDomain "go-clean-arc-sauna-shop-app/app/domain/error" //パッケージをエイリアス化
	"go-clean-arc-sauna-shop-app/pkg/ulid"
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

func newProduct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int) (*Product, error) {

	// ownerIDのバリデーション
	if !ulid.Isvalid(ownerID) {
		return nil, errDomain.NewError("オーナーIDの値が不正です 。")
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
