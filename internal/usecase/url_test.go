package usecase

import (
	"context"
	"github.com/Enthreeka/ozon-short-url/internal/apperror"
	"github.com/Enthreeka/ozon-short-url/internal/repo"
	"github.com/Enthreeka/ozon-short-url/internal/repo/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateShortUrl(t *testing.T) {
	var repo repo.URLRepository
	repo = mock.NewURLRepositoryMock()

	urlUsecase := NewURLUsecase(repo)
	urlSlice := []string{}

	tests := []struct {
		name    string
		wantErr error
	}{
		{
			name:    "https://yputube.com/video",
			wantErr: nil,
		},
		{
			name:    "https://www.ozon.ru/product/23-8-monitor-sanc-m2456h-chernyy-1221743240/?campaignId=286",
			wantErr: apperror.ErrURLExist,
		},
		{
			name:    "",
			wantErr: nil,
		},
		{
			name:    "https:/yputube.com/video",
			wantErr: nil,
		},
		{
			name:    "https//yputube.com/video",
			wantErr: apperror.ErrInvalidURL,
		},
		{
			name:    "https:yputube.com/video",
			wantErr: apperror.ErrInvalidURL,
		},
	}

	for _, tt := range tests {
		t.Run("create short url tests", func(t *testing.T) {
			shortUrl, err := urlUsecase.CreateShortUrl(context.Background(), tt.name)
			urlSlice = append(urlSlice, shortUrl)
			assert.Equal(t, tt.wantErr, err)
		})
	}

	for _, tt := range urlSlice {
		t.Run("get original url test", func(t *testing.T) {
			_, err := urlUsecase.GetOriginalUrl(context.Background(), tt)
			if err != nil {
				assert.Equal(t, apperror.ErrNotFound, err)
			}
		})
	}

}
