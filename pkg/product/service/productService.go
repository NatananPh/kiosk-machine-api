package service

import "github.com/NatananPh/kiosk-machine-api/pkg/product/model"

type ProductService interface {
	GetProducts() ([]model.Product, error)
	GetProductByID(id int) (model.Product, error)
	UpdateProduct(id int, product model.Product) (model.Product, error)
	DeleteProduct(id int) error
}