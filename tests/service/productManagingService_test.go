package tests

import (
	"testing"

	"github.com/NatananPh/kiosk-machine-api/entities"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/repository"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGettingProductByIDSuccess(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productEntity := &entities.Product{
		ID:     1,
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
	}

	productRepositoryMock.On("GetProductByID", 1).Return(productEntity, nil)

	product, err := productService.GetProductByID(1)
	assert.Nil(t, err)
	assert.Equal(t, productEntity.ID, product.ID)
	assert.Equal(t, productEntity.Name, product.Name)
	assert.Equal(t, productEntity.Price, product.Price)
	assert.Equal(t, productEntity.Amount, product.Amount)
}

func TestGettingProductByIDFailed(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productRepositoryMock.On("GetProductByID", 1).Return(&entities.Product{}, assert.AnError)

	product, err := productService.GetProductByID(1)
	assert.NotNil(t, err)
	assert.Equal(t, 0, product.ID)
	assert.Equal(t, "", product.Name)
	assert.Equal(t, uint(0), product.Price)
	assert.Equal(t, uint(0), product.Amount)
}

func TestGettingProductsSuccess(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productEntities := []*entities.Product{
		{
			ID:     1,
			Name:   "Product 1",
			Price:  100,
			Amount: 10,
			Category: "Category 1",
		},
		{
			ID:     2,
			Name:   "Product 2",
			Price:  200,
			Amount: 20,
			Category: "Category 2",
		},
	}

	productRepositoryMock.On("GetProducts", &model.ProductFilter{}).Return(productEntities, nil)
	products, err := productService.GetProducts(&model.ProductFilter{})
	assert.Nil(t, err)
	assert.Equal(t, len(productEntities), len(products))
	for i, product := range products {
		assert.Equal(t, productEntities[i].ID, product.ID)
		assert.Equal(t, productEntities[i].Name, product.Name)
		assert.Equal(t, productEntities[i].Price, product.Price)
		assert.Equal(t, productEntities[i].Amount, product.Amount)
	}
}

func TestGettingProductsFailed(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productRepositoryMock.On("GetProducts", &model.ProductFilter{}).Return([]*entities.Product{}, assert.AnError)
	products, err := productService.GetProducts(&model.ProductFilter{})
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestGettingProductsFilteredSuccess(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productEntities := []*entities.Product{
		{
			ID:     1,
			Name:   "Product 1",
			Price:  100,
			Amount: 10,
			Category: "Category 1",
		},
		{
			ID:     2,
			Name:   "Product 2",
			Price:  200,
			Amount: 20,
			Category: "Category 2",
		},
	}

	productRepositoryMock.On("GetProducts", &model.ProductFilter{Category: "Category 1"}).Return([]*entities.Product{productEntities[0]}, nil)
	products, err := productService.GetProducts(&model.ProductFilter{Category: "Category 1"})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(products))
	assert.Equal(t, productEntities[0].ID, products[0].ID)
	assert.Equal(t, productEntities[0].Name, products[0].Name)
	assert.Equal(t, productEntities[0].Price, products[0].Price)
	assert.Equal(t, productEntities[0].Amount, products[0].Amount)
}

func TestGettingProductsFilteredFailed(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productRepositoryMock.On("GetProducts", &model.ProductFilter{Category: "Category 1"}).Return([]*entities.Product{}, assert.AnError)
	products, err := productService.GetProducts(&model.ProductFilter{Category: "Category 1"})
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(products))
}

func TestCreatingProductSuccess(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productCreateRequest := &model.ProductCreateRequest{
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productEntity := &entities.Product{
		ID:     0,
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productRepositoryMock.On("CreateProduct", productEntity).Return(productEntity, nil)

	product, err := productService.CreateProduct(productCreateRequest)
	assert.Nil(t, err)
	assert.Equal(t, productEntity.ID, product.ID)
	assert.Equal(t, productEntity.Name, product.Name)
	assert.Equal(t, productEntity.Price, product.Price)
	assert.Equal(t, productEntity.Amount, product.Amount)
	assert.Equal(t, productEntity.Category, product.Category)
}

func TestCreatingProductFailed(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productCreateRequest := &model.ProductCreateRequest{
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productRepositoryMock.On("CreateProduct", mock.Anything).Return(&entities.Product{}, assert.AnError)

	product, err := productService.CreateProduct(productCreateRequest)
	assert.NotNil(t, err)
	assert.Equal(t, 0, product.ID)
	assert.Equal(t, "", product.Name)
	assert.Equal(t, uint(0), product.Price)
	assert.Equal(t, uint(0), product.Amount)
	assert.Equal(t, "", product.Category)
}

func TestUpdatingProductSuccess(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productUpdateRequest := &model.Product{
		ID:     1,
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productEntity := &entities.Product{
		ID:     1,
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productRepositoryMock.On("UpdateProduct", 1, productEntity).Return(productEntity, nil)

	product, err := productService.UpdateProduct(1, productUpdateRequest)
	assert.Nil(t, err)
	assert.Equal(t, productEntity.ID, product.ID)
	assert.Equal(t, productEntity.Name, product.Name)
	assert.Equal(t, productEntity.Price, product.Price)
	assert.Equal(t, productEntity.Amount, product.Amount)
	assert.Equal(t, productEntity.Category, product.Category)
}

func TestUpdatingProductFailed(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productUpdateRequest := &model.Product{
		ID:     1,
		Name:   "Product 1",
		Price:  100,
		Amount: 10,
		Category: "Category 1",
	}

	productRepositoryMock.On("UpdateProduct", 1, mock.Anything).Return(&entities.Product{}, assert.AnError)

	product, err := productService.UpdateProduct(1, productUpdateRequest)
	assert.NotNil(t, err)
	assert.Equal(t, 0, product.ID)
	assert.Equal(t, "", product.Name)
	assert.Equal(t, uint(0), product.Price)
	assert.Equal(t, uint(0), product.Amount)
	assert.Equal(t, "", product.Category)
}

func TestDeletingProductSuccess(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productRepositoryMock.On("DeleteProduct", 1).Return(nil)

	err := productService.DeleteProduct(1)
	assert.Nil(t, err)
}

func TestDeletingProductFailed(t *testing.T) {
	productRepositoryMock := new(repository.ProductRepositoryMock)
	productService := service.NewProductService(productRepositoryMock)

	productRepositoryMock.On("DeleteProduct", 1).Return(assert.AnError)

	err := productService.DeleteProduct(1)
	assert.NotNil(t, err)
}