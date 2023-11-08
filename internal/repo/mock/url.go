package mock

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/internal/entity"
	"github.com/Enthreeka/ozon-short-url/internal/repo"
	"github.com/jackc/pgx/v5"
)

type urlRepositoryMock struct {
	m map[string]string
}

func NewURLRepositoryMock() repo.URLRepository {
	m := map[string]string{
		"https://www.ozon.ru/product/23-8-monitor-sanc-m2456h-chernyy-1221743240/?campaignId=286": "https://www.ozon.ru/84lu0ycFda",
	}

	return &urlRepositoryMock{
		m,
	}
}

func (u *urlRepositoryMock) Create(ctx context.Context, url *entity.URL) error {
	u.m[url.ShortURL] = url.OriginalURL
	return nil
}

func (u *urlRepositoryMock) GetByShortURL(ctx context.Context, url string) (string, error) {
	originalURL, ok := u.m[url]
	if !ok {
		return "", pgx.ErrNoRows
	}

	return originalURL, nil
}

func (u *urlRepositoryMock) IsExistOriginalURL(ctx context.Context, url string) (bool, error) {
	for _, value := range u.m {
		if value == url {
			return true, nil
		}
	}

	return false, nil
}
