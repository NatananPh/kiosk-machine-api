package server

import (
	"github.com/NatananPh/kiosk-machine-api/pkg/auth/controller"
	"github.com/NatananPh/kiosk-machine-api/pkg/auth/repository"
	"github.com/NatananPh/kiosk-machine-api/pkg/auth/service"
)

func (s *echoServer) registerAuthRoutes() {
	authRepository := repository.NewAuthRepository(s.db)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	router := s.app.Group("/v1/auth")
	router.POST("/login", authController.Login)
}