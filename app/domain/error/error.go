package error

import "errors"

// errorインターフェースを暗黙的に実装(Errorメソッドをカスタム)
type Error struct {
	description string
}

func (e *Error) Error() string {
	return e.description
}

func NewError(s string) *Error {
	return &Error{
		description: s,
	}
}

var NotFoundErr = errors.New("not found")
