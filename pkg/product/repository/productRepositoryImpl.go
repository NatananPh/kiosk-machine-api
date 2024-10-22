package repository

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/exception"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"gorm.io/gorm"
)

type productRepositoryImpl struct{
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		db: db,
	}
}

func (p *productRepositoryImpl) CreateProduct(product *entities.Product) (*entities.Product, error) {
	if err := p.db.Create(&product).Error; err != nil {
		return &entities.Product{}, err
	}
	return product, nil
}

func (p *productRepositoryImpl) GetProducts(filter *model.ProductFilter) ([]*entities.Product, error) {
	var products []*entities.Product
	
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	if filter.Category != "" {
		var filteredProducts []*entities.Product
		for _, product := range products {
			if product.Category == filter.Category {
				filteredProducts = append(filteredProducts, product)
			}
		}
		return filteredProducts, nil
	}
	return products, nil
}

func (p *productRepositoryImpl) GetProductByID(id int) (*entities.Product, error) {
	var product *entities.Product

	if err := p.db.First(&product, id).Error; err != nil {
		return &entities.Product{}, err
	}
	return product, nil
}

func (p *productRepositoryImpl) UpdateProduct(id int, product *entities.Product) (*entities.Product, error) {
	result := p.db.Model(&entities.Product{}).Where("id = ?", id).Updates(product)
	if err := result.Error; err != nil {
		return &entities.Product{}, err
	}

	if result.RowsAffected == 0 {
		return &entities.Product{}, &exception.ProductNotFound{}
	}

	return product, nil
}

func (p *productRepositoryImpl) DeleteProduct(id int) error {
	result := p.db.Delete(&entities.Product{}, id)
	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return &exception.ProductNotFound{}
	}
	return nil
}

func (p *productRepositoryImpl) PurchaseProduct(id int) (*entities.Product, error) {
	var product *entities.Product

	if err := p.db.First(&product, id).Error; err != nil {
		return &entities.Product{}, err
	}

	if product.Amount == 0 {
		return &entities.Product{}, &exception.ProductPurchasing{}
	}
	product.Amount -= 1
	if err := p.db.Save(&product).Error; err != nil {
		return &entities.Product{}, err
	}
	return product, nil
}