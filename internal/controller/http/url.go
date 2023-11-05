package http

import (
	"context"
	"encoding/json"
	"github.com/Enthreeka/ozon-short-url/internal/apperror"
	"github.com/Enthreeka/ozon-short-url/internal/entity/dto"
	"github.com/Enthreeka/ozon-short-url/internal/usecase"
	"github.com/Enthreeka/ozon-short-url/pkg/logger"
	"net/http"
)

type urlServer struct {
	urlUsecase usecase.URLUsecase
	log        *logger.Logger
}

func NewUrlServerHandler(urlUsecase usecase.URLUsecase, log *logger.Logger) *urlServer {
	return &urlServer{
		urlUsecase: urlUsecase,
		log:        log,
	}
}

func (u *urlServer) CreateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var url dto.CreateURLRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&url)
	if err != nil {
		DecodingError(w)
		return
	}

	shortURL, err := u.urlUsecase.CreateShortUrl(context.Background(), url.OriginalURL)
	if err != nil {
		HandleError(w, err, apperror.ParseHTTPErrStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.Encode(map[string]string{
		"short_url": shortURL,
	})
}

func (u *urlServer) GetOriginalUrlHandler(w http.ResponseWriter, r *http.Request) {
	var url dto.GetURLRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&url)
	if err != nil {
		DecodingError(w)
		return
	}

	originalURL, err := u.urlUsecase.GetOriginalUrl(context.Background(), url.ShortURL)
	if err != nil {
		HandleError(w, err, apperror.ParseHTTPErrStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.Encode(map[string]string{
		"original_url": originalURL,
	})
}
