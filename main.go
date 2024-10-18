package main

import (
	"github.com/NatananPh/kiosk-machine-api/config"
	"github.com/NatananPh/kiosk-machine-api/database"
	"github.com/NatananPh/kiosk-machine-api/server"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(cfg.Database)
	echoServer := server.NewEchoServer(db.Connect(), cfg)
	echoServer.Start()
}