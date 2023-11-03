package app

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/internal/config"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
	"github.com/Enthreeka/ozon-short-url/pkg/postgres"
	"github.com/Enthreeka/ozon-short-url/pkg/redis"
)

func Run(log *logger.Logger, cfg *config.Config) error {
	// Connect to PostgreSQL
	psql, err := postgres.New(context.Background(), 5, cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}
	defer psql.Close()

	// Connect to Redis
	rds, err := redis.New(context.Background(), cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.MinIdleConns, cfg.Redis.DB)
	if err != nil {
		log.Error("redis is not working: %v", err)
	}
	defer rds.Close()

	return nil
}
