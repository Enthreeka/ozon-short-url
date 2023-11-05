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
	if r.Method == http.MethodPost {
		var url dto.CreateURLRequest
		d := json.NewDecoder(r.Body)
		err := d.Decode(&url)
		if err != nil {
			u.log.Error("json.NewDecoder: %v", err)
			DecodingError(w)
			return
		}

		shortURL, err := u.urlUsecase.CreateShortUrl(context.Background(), url.OriginalURL)
		if err != nil {
			u.log.Error("urlUsecase.CreateShortUrl: %v", err)
			HandleError(w, err, apperror.ParseHTTPErrStatusCode(err))
			return
		}

		w.WriteHeader(http.StatusCreated)
		e := json.NewEncoder(w)
		e.Encode(map[string]string{
			"short_url": shortURL,
		})
	}
}

func (u *urlServer) GetOriginalUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		shortURL := r.URL.Query().Get("shortURL")
		if shortURL == "" {
			u.log.Error("shortURL == zero value")
			QueryError(w)
			return
		}

		originalURL, err := u.urlUsecase.GetOriginalUrl(context.Background(), shortURL)
		if err != nil {
			u.log.Error("urlUsecase.GetOriginalUrl: %v", err)
			HandleError(w, err, apperror.ParseHTTPErrStatusCode(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		e := json.NewEncoder(w)
		e.Encode(map[string]string{
			"original_url": originalURL,
		})
	}
}
