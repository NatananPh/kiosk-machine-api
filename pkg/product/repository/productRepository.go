package repository

import "github.com/NatananPh/kiosk-machine-api/entities"

type ProductRepository interface {
	CreateProduct(product entities.Product) (entities.Product, error)
	GetProducts() ([]entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	UpdateProduct(id int, product entities.Product) (entities.Product, error)
	DeleteProduct(id int) error
}