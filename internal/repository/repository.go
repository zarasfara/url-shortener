package repository

import (
	"database/sql"

	"github.com/zarasfara/url-shortener/internal/repository/postgres"
)

type UrlShortenerStorage interface {
	SaveUrl(url, alias string) error
	GetURL(alias string) (string, error)
}

type Repository struct {
	UrlShortenerStorage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UrlShortenerStorage: postgres.NewUrlShortenerRepository(db),
	}
}
