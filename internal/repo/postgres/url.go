package postgres

import (
	"context"
	"errors"
	"github.com/Enthreeka/ozon-short-url/internal/entity"
	"github.com/Enthreeka/ozon-short-url/internal/repo"
	"github.com/Enthreeka/ozon-short-url/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type urlRepositoryPG struct {
	*postgres.Postgres
}

func NewURLRepositoryPG(pg *postgres.Postgres) repo.URLRepository {
	return &urlRepositoryPG{
		pg,
	}
}

func (u *urlRepositoryPG) Create(ctx context.Context, url *entity.URL) error {
	query := `INSERT INTO url (original_url,short_url) VALUES ($1,$2)`

	_, err := u.Pool.Exec(ctx, query, url.OriginalURL, url.ShortURL)
	return err
}

func (u *urlRepositoryPG) GetByShortURL(ctx context.Context, url string) (string, error) {
	query := `SELECT original_url FROM url WHERE short_url = $1`
	var originalURL string

	err := u.Pool.QueryRow(ctx, query, url).Scan(&originalURL)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", pgx.ErrNoRows
		}

		return "", err
	}

	return originalURL, nil
}

func (u *urlRepositoryPG) GetByOriginalURL(ctx context.Context, url string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1
				FROM url
				WHERE original_url = $1)`
	var exist bool

	err := u.Pool.QueryRow(ctx, query, url).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}
