package entities

type Role struct {
	ID   int    `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:varchar(255);not null"`
}