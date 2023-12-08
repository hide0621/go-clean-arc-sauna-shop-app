package cart

type cartProduct struct {
	productID string
	quantity  int
}

type Cart struct {
	userID   string
	products []cartProduct
}

func (cp *cartProduct) ProductID() string {
	return cp.productID
}

func (c *Cart) ProductIDs() []string {
	var productIDs []string
	for _, product := range c.products {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}

func (c *Cart) Products() []cartProduct {
	return c.products
}

func (cp *cartProduct) Quantity() int {
	return cp.quantity
}

func (c *Cart) UserID() string {
	return c.userID
}
