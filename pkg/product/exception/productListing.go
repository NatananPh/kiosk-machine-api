package exception

type ProductListing struct{}

func (p *ProductListing) Error() string {
	return "Product listing error"
}