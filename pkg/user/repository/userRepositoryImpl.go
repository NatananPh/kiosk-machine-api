package repository

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) GetUsers() ([]*entities.User, error) {
	var users []*entities.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}