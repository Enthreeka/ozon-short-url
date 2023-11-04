package repo

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/internal/entity"
)

type URLRepository interface {
	Create(ctx context.Context, url *entity.URL) error
	GetByShortURL(ctx context.Context, url string) (string, error)
	GetByOriginalURL(ctx context.Context, url string) (bool, error)
}
