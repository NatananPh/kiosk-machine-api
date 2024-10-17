package entities

import "time"

type Product struct {
	ID        int       `gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Price     uint      `gorm:"not null"`
	Amount    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}