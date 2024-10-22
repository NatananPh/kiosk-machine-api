package repository

import (
	"github.com/NatananPh/kiosk-machine-api/entities"
)

type AuthRepository interface {
	GetAuthUser(username string) (entities.User, error)
}