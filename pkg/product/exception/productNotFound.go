package exception

type ProductNotFound struct{}

func (p *ProductNotFound) Error() string {
	return "Product not found"
}