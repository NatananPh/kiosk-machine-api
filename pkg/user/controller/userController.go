package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	GetUsers(ctx echo.Context) error
}