package exception

type ProductOutOfStock struct{}

func (p *ProductOutOfStock) Error() string {
	return "Product out of stock"
}