package service

import "github.com/NatananPh/kiosk-machine-api/pkg/user/model"

type UserService interface {
	GetUsers() ([]model.User, error)
}