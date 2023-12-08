package ulid

import (
	"github.com/oklog/ulid/v2"
)

// ULIDの生成
func NewULID() string {
	return ulid.Make().String()
}

// エラーチェック機能（バリデーション機能）
func IsValid(s string) bool {
	_, err := ulid.Parse(s)
	return err == nil
}
