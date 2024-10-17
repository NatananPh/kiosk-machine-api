package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/NatananPh/kiosk-machine-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	once sync.Once
	db   *postgresDatabase
)

func NewPostgresDatabase(cfg *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode, cfg.Schema)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		log.Printf("Connected to database %s", cfg.DBName)
		db = &postgresDatabase{conn}

	})
	return db
}

func (db *postgresDatabase) Connect() *gorm.DB {
	return db.DB
}