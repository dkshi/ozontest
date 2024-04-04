package repository

type URL interface {
	CreateShortURL(originalURL, shortURL string) (string, error)
	GetOriginalURL(shortURL string) (string, error)
	StartGC()
}

type Repository struct {
	URL
}

func NewRepository() *Repository {
	return &Repository{URL: newURLInMemory()}
}
