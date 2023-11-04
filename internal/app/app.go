package app

import (
	"context"
	"fmt"
	"github.com/Enthreeka/ozon-short-url/internal/config"
	pb "github.com/Enthreeka/ozon-short-url/internal/controller/grpc"
	urlGrpc "github.com/Enthreeka/ozon-short-url/internal/controller/grpc"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
	"github.com/Enthreeka/ozon-short-url/pkg/postgres"
	"github.com/Enthreeka/ozon-short-url/pkg/redis"
	"google.golang.org/grpc"
	"net"
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

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCServer.Port))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	urlGrpcServer := urlGrpc.NewUrlSeverHandler(log)

	pb.RegisterUrlShortenServiceServer(s, urlGrpcServer)

	log.Info("Starting gRPC listener on port :" + cfg.GRPCServer.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
		return err
	}

	return nil
}
