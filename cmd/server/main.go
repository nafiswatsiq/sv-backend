package main

import (
	"log"
	"sharingvision/internal/app"
	"sharingvision/internal/config"
)

func main() {
	cfg := config.Load()
	server, err := app.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	if err := server.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
