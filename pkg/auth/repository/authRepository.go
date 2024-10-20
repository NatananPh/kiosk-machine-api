package repository

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
)

type AuthRepository interface {
	GetAuthUser(username string, password string) (entities.User, error)
}