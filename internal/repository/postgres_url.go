package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const urlsTable = "urls"

type URLPostgres struct {
	db *sqlx.DB
}

func NewURLPostgres(db *sqlx.DB) *URLPostgres {
	newURLPostgres := &URLPostgres{
		db: db,
	}
	newURLPostgres.CreateURLsTable()
	return newURLPostgres
}

// Day without migrations :)
func (r *URLPostgres) CreateURLsTable() {
	createQuery := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		url_id SERIAL PRIMARY KEY,
		original_url VARCHAR(300),
		short_url VARCHAR(50),
		ttl TIMESTAMP
	);`, urlsTable)
	_, err := r.db.Exec(createQuery) // I don't want to use MustExec because of panic, so if table was not create just print it out in logs
	if err != nil {
		logrus.Printf("error creating urls table: %s", err.Error())
	}
}

func (r *URLPostgres) CreateShortURL(originalURL, shortURL string) (string, error) {
	selectQuery := fmt.Sprintf("SELECT short_url, ttl FROM %s WHERE original_url=$1", urlsTable)
	row := r.db.QueryRow(selectQuery, originalURL)

	var returnedShortURL string
	var returnedTTL time.Time

	// If originalURL is already exists and it's ttl is not over
	if row.Scan(&returnedShortURL, &returnedTTL) == nil && returnedTTL.After(time.Now()) {
		return returnedShortURL, nil
	}

	insertQuery := fmt.Sprintf("INSERT INTO %s (original_url, short_url, ttl) VALUES ($1, $2, $3)", urlsTable)
	_, err := r.db.Exec(insertQuery, originalURL, shortURL, time.Now().Add(urlTTL))

	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (r *URLPostgres) GetOriginalURL(shortURL string) (string, error) {
	selectQuery := fmt.Sprintf("SELECT original_url, ttl FROM %s WHERE short_url=$1", urlsTable)
	row := r.db.QueryRow(selectQuery, shortURL)

	var returnedOriginalURL string
	var returnedTTL time.Time

	// If originalURL doesn't exist or it's ttl is over
	if row.Scan(&returnedOriginalURL, &returnedTTL) != nil || time.Now().After(returnedTTL) {
		return "", errors.New("error: there is no original URL for this short URL")
	}

	return returnedOriginalURL, nil
}

// Garbage collecting with cleanup inteval
func (r *URLPostgres) GC() {
	for {
		<-time.After(cleanupInterval)

		expiredURLs := r.getExpiredURLs()
		r.deleteExpiredURLs(expiredURLs)
	}
}

func (r *URLPostgres) StartGC() {
	go r.GC()
}

func (r *URLPostgres) getExpiredURLs() (urls []string) {
	selectQuery := fmt.Sprintf("SELECT original_url, ttl FROM %s", urlsTable)
	rows, err := r.db.Query(selectQuery)
	if err != nil {
		logrus.Printf("unable to get expired urls: %s", err.Error())
		return []string{}
	}

	for rows.Next() {
		var returnedOriginalURL string
		var returnedTTL time.Time

		err = rows.Scan(&returnedOriginalURL, &returnedTTL)
		if err != nil {
			logrus.Printf("unable to scan from rows: %s, skipping...", err.Error())
			continue
		}

		if time.Now().After(returnedTTL) {
			urls = append(urls, returnedOriginalURL)
		}
	}

	return urls
}

func (r *URLPostgres) deleteExpiredURLs(urls []string) {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE original_url=$1", urlsTable)
	for _, url := range urls {
		_, err := r.db.Exec(deleteQuery, url)
		if err != nil {
			logrus.Printf("unable to delete url: %s, skipping...", err.Error())
		}
	}
}
