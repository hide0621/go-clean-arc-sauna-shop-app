package ulid

import (
	"github.com/oklog/ulid/v2"
)

// エラーチェック機能（バリデーション機能）
func Isvalid(s string) bool {
	_, err := ulid.Parse(s)
	return err == nil
}
