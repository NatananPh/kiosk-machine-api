package controller

import (
	"github.com/NatananPh/kiosk-machine-api/pkg/user/service"
	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (uc *UserControllerImpl) GetUsers(ctx echo.Context) error {
	users, err := uc.userService.GetUsers()
	if err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, users)
}
