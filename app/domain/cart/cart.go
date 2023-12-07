package cart

type cartProduct struct {
	productID string
	quantity  int
}

type Cart struct {
	userID   string
	products []cartProduct
}
