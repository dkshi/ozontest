package service

import (
	"math/rand"
	"strings"

	"github.com/dkshi/shortener/internal/repository"
)

const (
	alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	lengthShortURL = 10
)

type URLService struct {
	repo *repository.Repository
}

func NewURLService(repo *repository.Repository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (s *URLService) CreateShortURL(originalURL string) (string, error) {
	return s.repo.CreateShortURL(originalURL, s.generateShortURL())
}

func (s *URLService) GetOriginalURL(shortURL string) (string, error) {
	return s.repo.GetOriginalURL(shortURL)
}

func (s *URLService) generateShortURL() string {
	// Generate random string with length 10
	shortURL := s.generateRandomString(lengthShortURL)
	_, err := s.repo.GetOriginalURL(shortURL)

	// Error not equals nil only if new shortURL doesn't exist
	// If shortURL already exists just make a new random string
	for err == nil {
		shortURL = s.generateRandomString(lengthShortURL)
		_, err = s.repo.GetOriginalURL(shortURL)
	}

	return shortURL
}

// Generates random string
func (s *URLService) generateRandomString(length int) string {
	sb := strings.Builder{}
	// Pick up random 10 bytes from alphabet
	for i := 0; i < length; i++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}
	return sb.String()
}
