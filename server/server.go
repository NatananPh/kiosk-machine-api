package server

import (
	"fmt"
	"sync"

	"github.com/NatananPh/kiosk-machine-api/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db *gorm.DB
	cfg *config.Config
}

var (
	once sync.Once
	server *echoServer
)

func NewEchoServer(db *gorm.DB, cfg *config.Config) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app: echoApp,
			db: db,
			cfg: cfg,
		}
	})
	return server
}

func (s *echoServer) Start() {
	s.app.GET("/health" , func(c echo.Context) error {
		return c.String(200, "OK")
	})
	s.httpListening()
}

func (s *echoServer) httpListening() {
	url := fmt.Sprintf(":%d",  s.cfg.Server.Port)
	if err := s.app.Start(url); err != nil {
		s.app.Logger.Fatalf("Failed to start server: %v", err)
	}
}