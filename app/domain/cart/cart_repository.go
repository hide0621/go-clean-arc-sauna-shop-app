package cart

import "context"

type CartRepository interface {
	FindByUserID(ctx context.Context, userID string) (*Cart, error)
	Save(ctx context.Context, cart *Cart) error
}
