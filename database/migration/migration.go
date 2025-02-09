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

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func migrate(tx *gorm.DB) error {
	migrations := []func(*gorm.DB) error{
		userMigration,
		productMigration,
		roleMigration,
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
