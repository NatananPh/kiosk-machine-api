package repository

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
)

type ProductRepository interface {
	CreateProduct(product entities.Product) (entities.Product, error)
	GetProducts(filter model.ProductFilter) ([]entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	UpdateProduct(id int, product entities.Product) (entities.Product, error)
	DeleteProduct(id int) error
	PurchaseProduct(id int) (entities.Product, error)
}