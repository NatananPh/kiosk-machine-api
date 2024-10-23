package entities

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement;not null"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	RoleID    int       `gorm:"not null"`
	Role      Role      `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
}

