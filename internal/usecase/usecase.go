package usecase

import (
	"context"
	pb "github.com/Enthreeka/ozon-short-url/internal/controller/grpc"
)

type URLUsecase interface {
	CreateShortUrl(context.Context, *pb.OriginalUrl) (*pb.ShortUrl, error)
	GetOriginalUrl(context.Context, *pb.ShortUrl) (*pb.OriginalUrl, error)
}
