package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/zarasfara/url-shortener/internal/database/postgres"
	repoErrors "github.com/zarasfara/url-shortener/internal/repository/errors"
)

type UrlShortenerRepository struct {
	db *sql.DB
}

func NewUrlShortenerRepository(db *sql.DB) *UrlShortenerRepository {
	return &UrlShortenerRepository{
		db: db,
	}
}

func (us *UrlShortenerRepository) SaveUrl(url, alias string) error {
	query := fmt.Sprintf("INSERT INTO %s (url, alias) VALUES ($1, $2)", postgres.UrlsTable)

	stmt, err := us.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(url, alias)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return repoErrors.ErrAliasAlreadyExists
			}
		}
		return err
	}

	return nil
}

func (us *UrlShortenerRepository) GetURL(alias string) (string, error) {
	const op = "tmp.sqlite.GetURL"

	stmt, err := us.db.Prepare("SELECT url FROM url WHERE alias = ?")
	if err != nil {
		return "", fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	var resURL string

	err = stmt.QueryRow(alias).Scan(&resURL)
	if err != nil {
		return "", fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return resURL, nil
}
