package app

import (
	"context"
	"fmt"
	"github.com/Enthreeka/ozon-short-url/internal/config"
	pb "github.com/Enthreeka/ozon-short-url/internal/controller/grpc"
	urlGrpc "github.com/Enthreeka/ozon-short-url/internal/controller/grpc"
	urlHttp "github.com/Enthreeka/ozon-short-url/internal/controller/http"
	rdsRepo "github.com/Enthreeka/ozon-short-url/internal/repo/redis"
	"github.com/Enthreeka/ozon-short-url/internal/usecase"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
	"github.com/Enthreeka/ozon-short-url/pkg/redis"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
)

func Run(log *logger.Logger, cfg *config.Config) error {
	if len(os.Args) < 2 {
		log.Fatal("There should be 2 launch flags")
	}

	// Connect to PostgreSQL
	//psql, err := postgres.New(context.Background(), 5, cfg.Postgres.URL)
	//if err != nil {
	//	log.Fatal("failed to connect PostgreSQL: %v", err)
	//}
	//defer psql.Close()

	// Connect to Redis
	rds, err := redis.New(context.Background(), cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.MinIdleConns, cfg.Redis.DB)
	if err != nil {
		log.Fatal("redis is not working: %v", err)
	}
	defer rds.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCServer.Port))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
		return err
	}

	//pgRepo "github.com/Enthreeka/ozon-short-url/internal/repo/postgres"
	//urlRepoPG := pgRepo.NewURLRepositoryPG(psql)
	urlRepoRDS := rdsRepo.NewURLRepositoryRedis(rds)
	urlUsecase := usecase.NewURLUsecase(urlRepoRDS, log)

	command := os.Args[1]
	switch command {
	case "grpc":
		s := grpc.NewServer()
		urlGRPCServer := urlGrpc.NewUrlSeverHandler(log, urlUsecase)

		pb.RegisterUrlShortenServiceServer(s, urlGRPCServer)

		log.Info("Starting gRPC listener on port :%s", cfg.GRPCServer.Port)
		if err := s.Serve(lis); err != nil {
			log.Fatal("failed to serve: %v", err)
			return err
		}
	case "http":
		urlHTTPServer := urlHttp.NewUrlServerHandler(urlUsecase, log)
		mux := http.NewServeMux()

		mux.HandleFunc("/api/create", urlHTTPServer.CreateShortUrlHandler)
		mux.HandleFunc("/api/get", urlHTTPServer.GetOriginalUrlHandler)

		log.Info("Starting http server on %s: %s%s", cfg.HTTTPServer.TypeServer, cfg.HTTTPServer.Hostname, cfg.HTTTPServer.Port)
		if err := http.ListenAndServe(cfg.HTTTPServer.Port, mux); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", err)
		}
	}

	return nil
}
