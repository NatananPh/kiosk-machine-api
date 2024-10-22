package service

import "github.com/NatananPh/kiosk-machine-api/pkg/product/model"

type ProductService interface {
	CreateProduct(product model.ProductCreateRequest) (model.Product, error)
	GetProducts(filter model.ProductFilter) ([]model.Product, error)
	GetProductByID(id int) (model.Product, error)
	UpdateProduct(id int, product model.Product) (model.Product, error)
	DeleteProduct(id int) error
	PurchaseProduct(id int, amount uint) (model.ProductPurchaseResponse, error)
}