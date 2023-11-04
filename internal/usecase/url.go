package usecase

import (
	"context"
	pb "github.com/Enthreeka/ozon-short-url/internal/controller/grpc"
)

type urlUsecase struct {
}

func NewURLUsecase() URLUsecase {
	return &urlUsecase{}
}

func (u *urlUsecase) CreateShortUrl(ctx context.Context, url *pb.OriginalUrl) (*pb.ShortUrl, error) {
	//TODO implement me
	panic("implement me")
}

func (u *urlUsecase) GetOriginalUrl(ctx context.Context, url *pb.ShortUrl) (*pb.OriginalUrl, error) {
	//TODO implement me
	panic("implement me")
}
