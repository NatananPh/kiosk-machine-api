package repository

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) CreateProduct(product *entities.Product) (*entities.Product, error) {
	args := m.Called(product)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *ProductRepositoryMock) GetProducts(filter *model.ProductFilter) ([]*entities.Product, error) {
	args := m.Called(filter)
	return args.Get(0).([]*entities.Product), args.Error(1)
}

func (m *ProductRepositoryMock) GetProductByID(id int) (*entities.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *ProductRepositoryMock) UpdateProduct(id int, product *entities.Product) (*entities.Product, error) {
	args := m.Called(id, product)
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *ProductRepositoryMock) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ProductRepositoryMock) PurchaseProduct(id int) (*entities.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Product), args.Error(1)
}