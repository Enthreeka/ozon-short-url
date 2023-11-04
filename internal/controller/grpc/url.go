package grpc

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
)

type urlServer struct {
	UnimplementedUrlShortenServiceServer

	log *logger.Logger
}

func NewUrlSeverHandler(log *logger.Logger) UrlShortenServiceServer {
	return &urlServer{
		log: log,
	}
}

func (u *urlServer) CreateShortUrl(ctx context.Context, req *OriginalUrl) (*ShortUrl, error) {

	panic("error")
}

func (u *urlServer) GetOriginalUrl(ctx context.Context, req *ShortUrl) (*OriginalUrl, error) {

	panic("error")

}
