package main

import (
	"github.com/Enthreeka/ozon-short-url/internal/config"
	"github.com/Enthreeka/ozon-short-url/internal/server"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
)

func main() {
	log := logger.New()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed load config: %v", err)
	}

	if err := server.Run(log, cfg); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
