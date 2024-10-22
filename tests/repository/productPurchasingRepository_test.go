package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/NatananPh/kiosk-machine-api/pkg/product/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestPurchaseProduct(t *testing.T) {
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
	mock.ExpectQuery("SELECT \\* FROM \"products\"").WithArgs(1, 1).WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"products\"").
		WithArgs("Product 1", 100, 9, "Category 1", sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	product, err := productRepository.PurchaseProduct(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, uint(100), product.Price)
	assert.Equal(t, uint(9), product.Amount)
	assert.Equal(t, "Category 1", product.Category)
}