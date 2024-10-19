package repository

import "github.com/NatananPh/kiosk-machine-api/entities"

type UserRepository interface {
	GetUsers() ([]entities.User, error)
}