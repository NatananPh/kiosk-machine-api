package controller

import (
	"github.com/NatananPh/kiosk-machine-api/pkg/auth/model"
	"github.com/NatananPh/kiosk-machine-api/pkg/auth/service"
	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}
	username := user.Username
	password := user.Password
	token, err := controller.AuthService.Login(username, password)
	if err != nil {
		return ctx.JSON(400, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, map[string]string{
		"token": token,
	})
}
