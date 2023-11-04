package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShorter(t *testing.T) {
	t.Run("return len and valid short url", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			shortUrl := GenerateShorterUrl()
			assert.Equal(t, 10, len(shortUrl))

		}
	})
}
