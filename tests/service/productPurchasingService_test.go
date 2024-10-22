package tests

import (
	"testing"

	"github.com/NatananPh/kiosk-machine-api/entities"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/repository"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/service"
	"github.com/stretchr/testify/assert"
)

func TestPurchasingProductSuccess(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productEntity := &entities.Product{
		ID:     1,
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
	}

	purchaseResponse := &model.ProductPurchaseResponse{
		ProductID: 1,
		Change:  map[string]int{
			"50": 1,
		},
	}

	productRepositoryMock.On("PurchaseProduct", 1).Return(productEntity, nil)
	productRepositoryMock.On("GetProductByID", 1).Return(productEntity, nil)

	product, err := productService.PurchaseProduct(1, 150)
	assert.Nil(t, err)
	assert.Equal(t, purchaseResponse.ProductID, productEntity.ID)
	assert.Equal(t, purchaseResponse.Change, product.Change)
}

func TestPurchasingProductInsufficientAmount(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productEntity := &entities.Product{
		ID:     1,
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
	}

	productRepositoryMock.On("PurchaseProduct", 1).Return(productEntity, assert.AnError)
	productRepositoryMock.On("GetProductByID", 1).Return(productEntity, nil)

	product, err := productService.PurchaseProduct(1, 50)
	assert.NotNil(t, err)
	assert.Equal(t, &model.ProductPurchaseResponse{}, product)
}

func TestPurchasingProductOutOfStock(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productEntity := &entities.Product{
		ID:     1,
		Name:   "Product 1",
		Price:  100,
		Amount: 0,
	}

	productRepositoryMock.On("PurchaseProduct", 1).Return(productEntity, nil)
	productRepositoryMock.On("GetProductByID", 1).Return(productEntity, nil)

	product, err := productService.PurchaseProduct(1, 100)
	assert.NotNil(t, err)
	assert.Equal(t, &model.ProductPurchaseResponse{}, product)
}