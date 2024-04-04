package repository

import (
	"errors"
	"sync"
	"time"
)

type urlItem struct {
	originalURL string
	expires     time.Time
}

type urlInMemoryDB struct {
	shortened map[string]urlItem // Shortened URLs
	originals map[string]string  // Key - originalURL, value - shortURL
}

type URLInMemory struct {
	sync.RWMutex
	db *urlInMemoryDB
}

func NewURLInMemory() *URLInMemory {
	return &URLInMemory{db: &urlInMemoryDB{
		shortened: make(map[string]urlItem),
		originals: make(map[string]string),
	}}
}

func (r *URLInMemory) CreateShortURL(originalURL, shortURL string) (string, error) {
	r.Lock()
	defer r.Unlock()

	// If originalURL already exists
	if shortURLFromOriginals, ok := r.db.originals[originalURL]; ok {
		// And originalURL was not expired, return shortURL
		if url := r.db.shortened[shortURLFromOriginals]; url.expires.After(time.Now()) {
			return shortURLFromOriginals, nil
		}
	}

	newURL := urlItem{
		originalURL: originalURL,
		expires:     time.Now().Add(urlTTL),
	}

	r.db.shortened[shortURL] = newURL
	r.db.originals[originalURL] = shortURL

	return shortURL, nil
}

func (r *URLInMemory) GetOriginalURL(shortURL string) (string, error) {
	url, ok := r.db.shortened[shortURL]
	if !ok || time.Now().After(url.expires) {
		return "", errors.New("error: there is no original URL for this short URL")
	}

	return url.originalURL, nil
}

// Garbage collecting with cleanup inteval
func (r *URLInMemory) GC() {
	for {
		<-time.After(cleanupInterval)

		expiredURLs := r.getExpiredURLs()
		r.deleteExpiredURLs(expiredURLs)
	}
}

func (r *URLInMemory) StartGC() {
	go r.GC()
}

func (r *URLInMemory) getExpiredURLs() (urls []string) {
	r.RLock()
	defer r.RUnlock()

	for k, v := range r.db.shortened {
		if time.Now().After(v.expires) {
			urls = append(urls, k)
		}
	}

	return urls
}

func (r *URLInMemory) deleteExpiredURLs(urls []string) {
	r.Lock()
	defer r.Unlock()

	for _, url := range urls {
		original := r.db.shortened[url]
		delete(r.db.shortened, url)
		delete(r.db.originals, original.originalURL)
	}
}
