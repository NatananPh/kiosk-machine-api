package service

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/repository"
)

type productServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productServiceImpl{
		productRepository: productRepository,
	}
}

func (p *productServiceImpl) GetProducts() ([]model.Product, error) {
	products, err := p.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	var productModels []model.Product
	for _, product := range products {
		productModels = append(productModels, model.Product{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
			Amount: product.Amount,
		})
	}

	return productModels, nil
}

func (p *productServiceImpl) GetProductByID(id int) (model.Product, error) {
	product, err := p.productRepository.GetProductByID(id)
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Amount: product.Amount,
	}, nil
}

func (p *productServiceImpl) UpdateProduct(id int, product model.Product) (model.Product, error) {
	productEntity := entities.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Amount: product.Amount,
	}

	updatedProductEntity, err := p.productRepository.UpdateProduct(id, productEntity)
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		ID:    updatedProductEntity.ID,
		Name:  updatedProductEntity.Name,
		Price: updatedProductEntity.Price,
		Amount: updatedProductEntity.Amount,
	}, nil
}

func (p *productServiceImpl) DeleteProduct(id int) error {
	return p.productRepository.DeleteProduct(id)
}