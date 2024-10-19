package server

import (
	"github.com/NatananPh/kiosk-machine-api/pkg/user/controller"
	"github.com/NatananPh/kiosk-machine-api/pkg/user/repository"
	"github.com/NatananPh/kiosk-machine-api/pkg/user/service"
)

func (s *echoServer) registerUserRoutes() {

	userRepository := repository.NewUserRepository(s.db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := s.app.Group("/v1/users")
	router.GET("", userController.GetUsers)
}