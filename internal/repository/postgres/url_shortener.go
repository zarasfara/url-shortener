package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/zarasfara/url-shortener/internal/database/postgres"
	"github.com/zarasfara/url-shortener/internal/logger/sl"
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
	const op = "repository.postgres.SaveUrl"
	query := fmt.Sprintf("INSERT INTO %s (url, alias) VALUES ($1, $2)", postgres.UrlsTable)

	// Выполнение запроса без явной подготовки
	_, err := us.db.Exec(query, url, alias)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				slog.Error("alias already exists", sl.WithError(repoErrors.ErrAliasAlreadyExists))
				return repoErrors.ErrAliasAlreadyExists
			}
		}
		slog.Error("failed to execute save URL statement", sl.WithError(err))
		return fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return nil
}

func (us *UrlShortenerRepository) GetUrl(alias string) (string, error) {
	const op = "repository.postgres.GetURL"

	stmt, err := us.db.Prepare("SELECT url FROM urls WHERE alias = $1")
	if err != nil {
		slog.Error("Failed to prepare statement", sl.WithError(err))
		return "", fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	defer stmt.Close()

	var url string
	err = stmt.QueryRow(alias).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("No rows found for alias", sl.WithError(err))
			return "", fmt.Errorf("%s: no rows found: %w", op, err)
		}
		slog.Error("Failed to execute get URL statement", sl.WithError(err))
		return "", fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return url, nil
}
