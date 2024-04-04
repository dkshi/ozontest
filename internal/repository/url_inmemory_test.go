package repository

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLInMemory_CreateShortURL(t *testing.T) {
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

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			uim := NewURLInMemory()
			shortURL, err := uim.CreateShortURL(testCase.inputOriginalURL, testCase.inpurtShortURL)
			assert.Equal(t, testCase.expectedShortURL, shortURL)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestURLInMemory_GetOriginalURL(t *testing.T) {
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

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			uim := NewURLInMemory()
			if testCase.expectedOriginalURL != "" {
				uim.CreateShortURL(testCase.inputOriginalURL, testCase.inputShortURL)
			}

			originalURL, err := uim.GetOriginalURL(testCase.inputShortURL)
			assert.Equal(t, testCase.expectedOriginalURL, originalURL)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
