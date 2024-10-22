package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NatananPh/kiosk-machine-api/entities"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	productRepository := repository.NewProductRepository(gormDB)
	
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"products\"").
		WithArgs("Product 1", 100, 10, "Category 1", time.Now(), time.Now()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	product, err := productRepository.CreateProduct(&entities.Product{
		Name:     "Product 1",
		Price:    100,
		Amount:   10,
		Category: "Category 1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, uint(100), product.Price)
	assert.Equal(t, uint(10), product.Amount)
	assert.Equal(t, "Category 1", product.Category)
}

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	productRepository := repository.NewProductRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name", "price", "amount", "category", "created_at", "updated_at"}).
		AddRow(1, "Product 1", 100, 10, "Category 1", time.Now(), time.Now()).
		AddRow(2, "Product 2", 200, 20, "Category 2", time.Now(), time.Now())
	mock.ExpectQuery("SELECT \\* FROM \"products\"").WillReturnRows(rows)

	products, err := productRepository.GetProducts(&model.ProductFilter{})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(products))

	assert.Equal(t, 1, products[0].ID)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, uint(100), products[0].Price)
	assert.Equal(t, uint(10), products[0].Amount)
	assert.Equal(t, "Category 1", products[0].Category)

	assert.Equal(t, 2, products[1].ID)
	assert.Equal(t, "Product 2", products[1].Name)
	assert.Equal(t, uint(200), products[1].Price)
	assert.Equal(t, uint(20), products[1].Amount)
	assert.Equal(t, "Category 2", products[1].Category)
}

func TestGetProductByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	productRepository := repository.NewProductRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name", "price", "amount", "category", "created_at", "updated_at"}).
		AddRow(1, "Product 1", 100, 10, "Category 1", time.Now(), time.Now())
	mock.ExpectQuery("SELECT \\* FROM \"products\"").WithArgs(1,1).WillReturnRows(rows)

	product, err := productRepository.GetProductByID(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, uint(100), product.Price)
	assert.Equal(t, uint(10), product.Amount)
	assert.Equal(t, "Category 1", product.Category)
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	productRepository := repository.NewProductRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"products\"").
		WithArgs(1,"Updated Product", 150, 15, "Updated Category", sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	
	product := &entities.Product{
		ID:       1,
		Name:     "Updated Product",
		Price:    150,
		Amount:   15,
		Category: "Updated Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	product, err = productRepository.UpdateProduct(1, product)
	assert.Nil(t, err)
	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Updated Product", product.Name)
	assert.Equal(t, uint(150), product.Price)
	assert.Equal(t, uint(15), product.Amount)
	assert.Equal(t, "Updated Category", product.Category)
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	productRepository := repository.NewProductRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM \"products\" WHERE \"products\".\"id\" = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = productRepository.DeleteProduct(1)
	assert.Nil(t, err)
}
