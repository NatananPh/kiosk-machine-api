package service

import "github.com/NatananPh/kiosk-machine-api/pkg/product/repository"

type productServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductServiceImpl(productRepository repository.ProductRepository) ProductService {
	return &productServiceImpl{
		productRepository: productRepository,
	}
}