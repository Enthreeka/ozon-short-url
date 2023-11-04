package usecase

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/internal/apperror"
	"github.com/Enthreeka/ozon-short-url/internal/entity"
	"github.com/Enthreeka/ozon-short-url/internal/repo"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
)

type urlUsecase struct {
	urlRepo repo.URLRepository

	log *logger.Logger
}

func NewURLUsecase(urlRepo repo.URLRepository, log *logger.Logger) URLUsecase {
	return &urlUsecase{
		urlRepo: urlRepo,
		log:     log,
	}
}

func (u *urlUsecase) CreateShortUrl(ctx context.Context, originalURL string) (string, error) {
	exist, err := u.urlRepo.GetByOriginalURL(ctx, originalURL)
	if err != nil {
		return "", apperror.NewError("Failed to check original url", err)
	}

	if exist {
		return "", apperror.ErrURLExist
	}

	identifier := GenerateShorterUrl()

	shortUrl, err := PrependBaseURL(originalURL, identifier)
	if err != nil {
		return "", apperror.ErrInvalidURL
	}

	url := &entity.URL{
		OriginalURL: originalURL,
		ShortURL:    shortUrl,
	}

	err = u.urlRepo.Create(ctx, url)
	if err != nil {
		return "", apperror.NewError("Failed to create short url", err)
	}

	return shortUrl, nil
}

func (u *urlUsecase) GetOriginalUrl(ctx context.Context, shortURL string) (string, error) {
	originalURL, err := u.urlRepo.GetByShortURL(ctx, shortURL)
	if err != nil {
		return "", apperror.ErrNotFound
	}

	return originalURL, nil
}