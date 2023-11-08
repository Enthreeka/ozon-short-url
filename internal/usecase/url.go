package usecase

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/internal/apperror"
	"github.com/Enthreeka/ozon-short-url/internal/entity"
	"github.com/Enthreeka/ozon-short-url/internal/repo"
	"github.com/google/uuid"
)

type urlUsecase struct {
	urlRepo repo.URLRepository
}

func NewURLUsecase(urlRepo repo.URLRepository) URLUsecase {
	return &urlUsecase{
		urlRepo: urlRepo,
	}
}

func (u *urlUsecase) CreateShortUrl(ctx context.Context, originalURL string) (string, error) {
	if !entity.IsUrl(originalURL) {
		return "", apperror.ErrInvalidURL
	}

	exist, err := u.urlRepo.IsExistOriginalURL(ctx, originalURL)
	if err != nil {
		return "", apperror.NewError("Failed to check original url", err)
	}

	if exist {
		return "", apperror.ErrURLExist
	}

	identifier := GenerateShorterUrl()

	shortURL, err := PrependBaseURL(originalURL, identifier)
	if err != nil {
		return "", apperror.ErrInvalidURL
	}

	url := &entity.URL{
		ID:          uuid.New().String(),
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	}

	err = u.urlRepo.Create(ctx, url)
	if err != nil {
		return "", apperror.NewError("Failed to create short url", err)
	}

	return shortURL, nil
}

func (u *urlUsecase) GetOriginalUrl(ctx context.Context, shortURL string) (string, error) {
	originalURL, err := u.urlRepo.GetByShortURL(ctx, shortURL)
	if err != nil {
		return "", apperror.ErrNotFound
	}

	return originalURL, nil
}
