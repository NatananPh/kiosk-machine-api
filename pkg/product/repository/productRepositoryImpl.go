package repository

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
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

func (p *productRepositoryImpl) CreateProduct(product entities.Product) (entities.Product, error) {
	if err := p.db.Create(&product).Error; err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (p *productRepositoryImpl) GetProducts() ([]entities.Product, error) {
	var products []entities.Product
	
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepositoryImpl) GetProductByID(id int) (entities.Product, error) {
	var product entities.Product

	if err := p.db.First(&product, id).Error; err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (p *productRepositoryImpl) UpdateProduct(id int, product entities.Product) (entities.Product, error) {
	if err := p.db.Model(&entities.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (p *productRepositoryImpl) DeleteProduct(id int) error {
	if err := p.db.Delete(&entities.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}