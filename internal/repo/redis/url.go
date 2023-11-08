package redis

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/internal/entity"
	"github.com/Enthreeka/ozon-short-url/internal/repo"
	rds "github.com/Enthreeka/ozon-short-url/pkg/redis"
	"time"
)

type urlRepositoryRedis struct {
	*rds.Redis
}

func NewURLRepositoryRedis(redis *rds.Redis) repo.URLRepository {
	return &urlRepositoryRedis{
		redis,
	}
}

func (u *urlRepositoryRedis) Create(ctx context.Context, url *entity.URL) error {
	err := u.Rds.Set(ctx, url.OriginalURL, url.ShortURL, 360*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u *urlRepositoryRedis) GetByShortURL(ctx context.Context, url string) (string, error) {
	iter := u.Rds.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		val := u.Rds.Get(ctx, iter.Val()).Val()
		if val == url {
			return iter.Val(), nil
		}
	}

	if err := iter.Err(); err != nil {
		return "", err
	}

	return "", nil
}

func (u *urlRepositoryRedis) IsExistOriginalURL(ctx context.Context, url string) (bool, error) {
	exists, err := u.Rds.Exists(ctx, url).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}
