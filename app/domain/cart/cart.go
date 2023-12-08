package cart

type cartProduct struct {
	productID string
	quantity  int
}

type Cart struct {
	userID   string
	products []cartProduct
}

func (c *Cart) ProductIDs() []string {
	var productIDs []string
	for _, product := range c.products {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}
