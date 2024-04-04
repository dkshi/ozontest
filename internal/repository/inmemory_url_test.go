package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLInMemory_CreateShortURL(t *testing.T) {
	uim := newURLInMemory()
	uim.StartGC()

	testTable := []struct {
		name             string
		inputOriginalURL string
		inpurtShortURL   string
		expectedShortURL string
		expectedError    error
	}{
		{
			name:             "OK",
			inputOriginalURL: "https://vk.com/dkshii",
			inpurtShortURL:   "sh.ort/9218jms",
			expectedShortURL: "sh.ort/9218jms",
			expectedError:    nil,
		},
	}

	for _, test := range testTable {
		shortURL, err := uim.CreateShortURL(test.inputOriginalURL, test.inpurtShortURL)
		assert.Equal(t, test.expectedShortURL, shortURL)
		assert.Equal(t, test.expectedError, err)
	}
}

func TestURLInMemory_GetOriginalURL(t *testing.T) {
	uim := newURLInMemory()
	uim.StartGC()

	testTable := []struct {
		name                string
		inputOriginalURL    string
		inputShortURL       string
		expectedOriginalURL string
		expectedError       error
	}{
		{
			name:                "OK",
			inputShortURL:       "9218jms",
			inputOriginalURL:    "https://vk.com/dkshii",
			expectedOriginalURL: "https://vk.com/dkshii",
			expectedError:       nil,
		},
		{
			name:                "No original URL",
			inputOriginalURL:    "https://error.error",
			inputShortURL:       "92128jms",
			expectedOriginalURL: "",
			expectedError:       errors.New("error: there is no original URL for this short URL"),
		},
	}

	for _, test := range testTable {
		if test.expectedOriginalURL != "" {
			uim.CreateShortURL(test.inputOriginalURL, test.inputShortURL)
		}

		originalURL, err := uim.GetOriginalURL(test.inputShortURL)
		assert.Equal(t, test.expectedOriginalURL, originalURL)
		assert.Equal(t, test.expectedError, err)
	}
}
