package serverapi

import (
	"testing"

	shortenerv1 "github.com/dkshi/shortener/gen"
	"github.com/dkshi/shortener/internal/service"
	mock_service "github.com/dkshi/shortener/internal/service/mocks"
	"github.com/golang/mock/gomock"
)

func TestServerAPI_CreateShortURL(t *testing.T) {
	type mockBehavior func(s *mock_service.MockURL, originalURL string)

	testTable := []struct {
		name                   string
		inputOriginalURL       *shortenerv1.OriginalURL
		inputStringOriginalURL string
		mockBehavior           mockBehavior
		expectedShortURL       *shortenerv1.ShortURL
		expectedError          error
	}{
		{
			name: "OK",
			inputOriginalURL: &shortenerv1.OriginalURL{
				Url: "http://vk.com/dkshii",				
			},
			inputStringOriginalURL: "http://vk.com/dkshii",
			mockBehavior: func(s *mock_service.MockURL, originalURL string) {
				s.EXPECT().CreateShortURL(originalURL).Return("abcDeFghki", nil)
			},
			expectedShortURL: &shortenerv1.ShortURL{
				Url: "abcDeFghki",
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testTable {
		c := gomock.NewController(t)
		defer c.Finish()

		urlMockService := mock_service.NewMockURL(c)
		testCase.mockBehavior(urlMockService, testCase.inputStringOriginalURL)

		s := &service.Service{URL: urlMockService}
		
		// допиши тест чертовка
	}
}
