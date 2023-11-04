package main

import (
	"github.com/Enthreeka/ozon-short-url/internal/app"
	"github.com/Enthreeka/ozon-short-url/internal/config"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
)

func main() {
	log := logger.New()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("%failed load config: %v", err)
	}

	if err := app.Run(log, cfg); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
