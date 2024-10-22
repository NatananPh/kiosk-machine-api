package model

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Amount   uint   `json:"amount"`
	Category string `json:"category"`
}
type ProductCreateRequest struct {
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Amount   uint   `json:"amount"`
	Category string `json:"category"`
}

type ProductFilter struct {
	Category string `query:"category" validate:"omitempty,max=64"`
}

type ProductPurchaseRequest struct {
	PaymentAmount uint `json:"payment_amount"`
}

type ProductPurchaseResponse struct {
	ProductID int            `json:"product_id"`
	Change    map[string]int `json:"change"`
}