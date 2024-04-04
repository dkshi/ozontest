package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestURLPostgres_CreateShortURL(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("error creating mock db: %s", err.Error())
	}
	defer db.Close()

	r := NewURLPostgres(db)

	testTable := []struct {
		name             string
		inputOriginalURL string
		inputShortURL    string
		expectedResult   string
		expectedError    error
	}{
		{
			name:             "OK",
			inputOriginalURL: "http://ozon.ru",
			inputShortURL:    "abcdefghij",
			expectedResult:   "abcdefghij",
			expectedError:    nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mock.ExpectQuery("SELECT short_url, ttl FROM urls WHERE original_url=?").
				WithArgs(testCase.inputOriginalURL).
				WillReturnError(errors.New("already exists"))

			mock.ExpectExec("INSERT INTO urls").
				WithArgs(testCase.inputOriginalURL, testCase.inputShortURL, sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(0, 1))

			result, err := r.CreateShortURL(testCase.inputOriginalURL, testCase.inputShortURL)

			assert.Equal(t, result, testCase.expectedResult)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}

func TestURLPostgres_GetOriginalURL(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("error creating mock db: %s", err.Error())
	}
	defer db.Close()

	r := NewURLPostgres(db)

	tests := []struct {
		name           string
		inputShortURL  string
		expectedResult string
		expectedError  error
	}{
		{
			name:           "OK",
			inputShortURL:  "abcdefghij",
			expectedResult: "http://ozon.ru",
			expectedError:  nil,
		},
		{
			name:           "No original URL",
			inputShortURL:  "abcdefghij",
			expectedResult: "",
			expectedError:  errors.New("error: there is no original URL for this short URL"),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			rows := sqlmock.NewRows([]string{"original_url", "ttl"})

			if testCase.expectedError == nil {
				rows = rows.AddRow(testCase.expectedResult, time.Now().Add(time.Hour))
			}

			mock.ExpectQuery("SELECT original_url, ttl FROM urls WHERE short_url=?").
				WithArgs(testCase.inputShortURL).
				WillReturnRows(rows)

			result, err := r.GetOriginalURL(testCase.inputShortURL)

			assert.Equal(t, result, testCase.expectedResult)
			assert.Equal(t, err, testCase.expectedError)
		})
	}
}
