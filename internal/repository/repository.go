package repository

import "database/sql"

type Repository struct {
	UrlShortenerStorage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UrlShortenerStorage: NewUrlShortenerRepository(db),
	}
}
