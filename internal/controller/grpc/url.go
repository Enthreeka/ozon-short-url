package grpc

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/internal/apperror"
	"github.com/Enthreeka/ozon-short-url/internal/usecase"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
	"google.golang.org/grpc/status"
)

type urlServer struct {
	log        *logger.Logger
	urlUsecase usecase.URLUsecase

	UnimplementedUrlShortenServiceServer
}

func NewUrlSeverHandler(log *logger.Logger, urlUsecase usecase.URLUsecase) UrlShortenServiceServer {
	return &urlServer{
		log:        log,
		urlUsecase: urlUsecase,
	}
}

func (u *urlServer) CreateShortUrl(ctx context.Context, req *OriginalUrlRequest) (*ShortUrlResponse, error) {
	shortURL, err := u.urlUsecase.CreateShortUrl(ctx, req.URL)
	if err != nil {
		u.log.Error("urlUsecase.CreateShortUrl: %v", err)
		return nil, status.Errorf(apperror.ParseGRPCErrStatusCode(err), "CreateShortUrl: %v", err)
	}

	return &ShortUrlResponse{URL: shortURL}, nil
}

func (u *urlServer) GetOriginalUrl(ctx context.Context, req *ShortUrlRequest) (*OriginalUrlResponse, error) {
	originalURL, err := u.urlUsecase.GetOriginalUrl(ctx, req.URL)
	if err != nil {
		u.log.Error("urlUsecase.GetOriginalUrl: %v", err)
		return nil, status.Errorf(apperror.ParseGRPCErrStatusCode(err), "GetOriginalUrl: %v", err)
	}

	return &OriginalUrlResponse{URL: originalURL}, nil
}
