package usecase

import (
	"context"
)

type URLUsecase interface {
	CreateShortUrl(ctx context.Context, originalURL string) (string, error)
	GetOriginalUrl(ctx context.Context, shortURL string) (string, error)
}
