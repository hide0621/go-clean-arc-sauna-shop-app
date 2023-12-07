package order

import "time"

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
