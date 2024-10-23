package main

import (
	"fmt"

	"github.com/NatananPh/kiosk-machine-api/config"
	"github.com/NatananPh/kiosk-machine-api/database"
	"github.com/NatananPh/kiosk-machine-api/entities"
	"gorm.io/gorm"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg.Database)
	fmt.Println(db.Connect())

	tx := db.Connect().Begin()
	migrate(tx)
	addRole(tx)
	addUser(tx)
	addProduct(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func addRole(tx *gorm.DB) error {
	var roles = []entities.Role{
		{
			Name: "Admin",
		},
		{
			Name: "User",
		},
	}
	tx.CreateInBatches(roles, len(roles))
	return nil
}

func addUser(tx *gorm.DB) error {
	var users = []entities.User{
		{
			Username: "admin",
			Password: "$2a$10$Tou3DwmFhYmnYj/CUAiBoeUxxUcy5mK150PXMY7Kj02x2PWAgQDye",
			RoleID:   1,
		},
		{
			Username: "user",
			Password: "$2a$10$f.wBU4dX4hSteATgSrSr7uYnP4LEoM6Ut8iUotK0Mx79OkLlfVOtK",
			RoleID:   2,
		},
	}
	tx.CreateInBatches(users, len(users))
	return nil
}

func addProduct(tx *gorm.DB) error {
	var products = []entities.Product{
		{
			Name:   "Coca Cola",
			Price:  20,
			Amount: 10,
			Category: "Drink",
		},
		{
			Name:   "Pepsi",
			Price:  20,
			Amount: 10,
			Category: "Drink",
		},
		{
			Name:   "Fanta",
			Price:  20,
			Amount: 10,
			Category: "Drink",
		},
		{
			Name:   "Sprite",
			Price:  20,
			Amount: 10,
			Category: "Drink",
		},	
		{
			Name:   "Lay",
			Price:  20,
			Amount: 10,
			Category: "Snack",
		},
		{
			Name:   "Doritos",
			Price:  20,
			Amount: 10,
			Category: "Snack",
		},
		{
			Name:   "TestOutOfStock",
			Price:  20,
			Amount: 0,
			Category: "Snack",
		},
	}
	tx.CreateInBatches(products, len(products))
	return nil
}

func migrate(tx *gorm.DB) error {
	migrations := []func(*gorm.DB) error{
		roleMigration,
		userMigration,
		productMigration,
	}

	for _, migration := range migrations {
		if err := migration(tx); err != nil {
			return err
		}
	}
	return nil
}

func userMigration(tx *gorm.DB) error {
	return tx.Migrator().CreateTable(&entities.User{})
}

func productMigration(tx *gorm.DB) error {
	return tx.Migrator().CreateTable(&entities.Product{})
}

func roleMigration(tx *gorm.DB) error {
	return tx.Migrator().CreateTable(&entities.Role{})
}
