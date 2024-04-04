package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	urlTTL          = time.Hour * 24 * 7 // Time to live of URL, 7 days
	cleanupInterval = time.Hour * 24     // Interval of cleaning expired keys, 1
)

type URL interface {
	CreateShortURL(originalURL, shortURL string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
	StartGC()
}

type Repository struct {
	URL
}

func NewRepository(db *sqlx.DB) *Repository {
	if db != nil {
		return &Repository{URL: NewURLPostgres(db)}
	}
	return &Repository{URL: NewURLInMemory()}
}
