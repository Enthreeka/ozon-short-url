package usecase

import (
	"github.com/google/uuid"
	"strings"
)

// All possible characters in the short link
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

// Length of the charset constant
const alphabetLen = uint32(len(charset))

// GenerateShorterUrl generates a short link
func GenerateShorterUrl() string {
	// To get rid of unnecessary allocations, we use strings.Builder
	var builder strings.Builder
	// This slice storage 10 random index
	indexCharsetArr := make([]uint32, 0, 10)

	// Generate unique random digit with uuid
	randomDigit := uuid.New().ID()

	// Loop to create 10 random indexes
	for len(indexCharsetArr) < 10 {
		// Append the remainder of the division in slice indexCharsetArr
		indexCharsetArr = append(indexCharsetArr, randomDigit%alphabetLen)
		// Reducing randomDigit by one number
		randomDigit /= 10
	}

	// Loop for taking value from charset by index
	for _, valueIndex := range indexCharsetArr {
		// We add all the resulting values in strings.Builder
		builder.WriteString(string(charset[valueIndex]))
	}

	// Return short url
	return builder.String()
}
