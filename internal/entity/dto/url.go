package dto

type CreateURLRequest struct {
	OriginalURL string `json:"original_url"`
}

type GetURLRequest struct {
	ShortURL string `json:"short_url"`
}
