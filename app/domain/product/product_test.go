package product

import (
	"go-clean-arc-sauna-shop-app/pkg/ulid"
	"testing"
)

func TestNewProduct(t *testing.T) {

	ownerID := ulid.NewULID()

	type args struct {
		ownerID     string
		name        string
		description string
		price       int64
		stock       int
	}

	tests := []struct {
		name    string
		args    args
		want    *Product
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			want: &Product{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			wantErr: false,
		},
		{
			name: "異常系: オーナーIDが不正",
			args: args{
				ownerID:     "test",
				name:        "test",
				description: "test",
				price:       100,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 名前が不正",
			args: args{
				ownerID:     ownerID,
				name:        "",
				description: "test",
				price:       100,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 商品説明が不正",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "",
				price:       100,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系: 価格が不正",
			args: args{
				ownerID:     ownerID,
				name:        "test",
				description: "test",
				price:       0,
				stock:       100,
			},
			want:    nil,
			wantErr: true,
		},
	}

}
