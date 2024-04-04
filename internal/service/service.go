package service

import (
	"github.com/dkshi/shortener/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type URL interface {
	CreateShortURL(originalURL string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
}

type Service struct {
	URL
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		URL: NewURLService(repo),
	}
}
