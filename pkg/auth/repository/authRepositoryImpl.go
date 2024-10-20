package repository

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (a *AuthRepositoryImpl) GetAuthUser(username string, password string) (entities.User, error) {
	var user entities.User

	if err := a.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}