package service

import (
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (m *ProductServiceMock) CreateProduct(product *model.ProductCreateRequest) (*model.Product, error) {
	ret := m.Called(product)

	r0 := ret.Get(0)
	if r0 == nil {
		return nil, ret.Error(1)
	}
	return r0.(*model.Product), ret.Error(1)
}

func (m *ProductServiceMock) GetProducts(filter *model.ProductFilter) ([]*model.Product, error) {
	ret := m.Called(filter)

	r0 := ret.Get(0)
	if r0 == nil {
		return nil, ret.Error(1)
	}
	return r0.([]*model.Product), ret.Error(1)
}

func (m *ProductServiceMock) GetProductByID(id int) (*model.Product, error) {
	ret := m.Called(id)

	r0 := ret.Get(0)
	if r0 == nil {
		return nil, ret.Error(1)
	}
	return r0.(*model.Product), ret.Error(1)
}

func (m *ProductServiceMock) UpdateProduct(id int, product *model.Product) (*model.Product, error) {
	ret := m.Called(id, product)

	r0 := ret.Get(0)
	if r0 == nil {
		return nil, ret.Error(1)
	}
	return r0.(*model.Product), ret.Error(1)
}

func (m *ProductServiceMock) DeleteProduct(id int) error {
	ret := m.Called(id)

	return ret.Error(0)
}

func (m *ProductServiceMock) PurchaseProduct(id int, amount uint) (*model.ProductPurchaseResponse, error) {
	ret := m.Called(id, amount)

	r0 := ret.Get(0)
	if r0 == nil {
		return nil, ret.Error(1)
	}
	return r0.(*model.ProductPurchaseResponse), ret.Error(1)
}