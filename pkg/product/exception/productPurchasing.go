package exception

type ProductPurchasing struct{}

func (p *ProductPurchasing) Error() string {
	return "Product Out of Stock"
}