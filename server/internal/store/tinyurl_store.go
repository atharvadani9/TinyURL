package store

import (
	"database/sql"
)

type TinyURL struct {
	ID          int    `json:"id"`
	OriginalURL string `json:"originalUrl"`
	ShortURL    string `json:"shortUrl"`
	CreatedAt   string `json:"createdAt"`
}

type PostgresTinyURLStore struct {
	db *sql.DB
}

func NewPostgresTinyURLStore(db *sql.DB) *PostgresTinyURLStore {
	return &PostgresTinyURLStore{db: db}
}

type TinyURLStore interface {
	Create(tinyURL *TinyURL) error
	Get(shortURL string) (*TinyURL, error)
}

func (s *PostgresTinyURLStore) Create(tinyURL *TinyURL) error {
	query := `
		INSERT INTO tinyurl (original_url, short_url)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	err := s.db.QueryRow(query, tinyURL.OriginalURL, tinyURL.ShortURL).Scan(&tinyURL.ID, &tinyURL.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresTinyURLStore) Get(shortURL string) (*TinyURL, error) {
	query := `
		SELECT id, original_url, created_at
		FROM tinyurl
		WHERE short_url = $1
	`

	tinyURL := &TinyURL{}
	err := s.db.QueryRow(query, shortURL).Scan(&tinyURL.ID, &tinyURL.OriginalURL, &tinyURL.CreatedAt)
	if err != nil {
		return nil, err
	}

	return tinyURL, nil
}
