package repository

import "database/sql"

type UrlShortenerStorage interface {
	HelloWorld() string
}

type UrlShortener struct {
	db *sql.DB
}

func NewUrlShortener(db *sql.DB) *UrlShortener {
	return &UrlShortener{
		db: db,
	}
}

func (us UrlShortener) HelloWorld() string {
	panic("implement me")
}