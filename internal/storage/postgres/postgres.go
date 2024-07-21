package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/zarasfara/url-shortener/internal/config"
	"github.com/zarasfara/url-shortener/internal/logger/sl"
)

func New(cfg config.Config, logger *slog.Logger) *sql.DB {

	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DB.Username, cfg.DB.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.DB.Host, cfg.DB.Port),
		Path:   "/" + cfg.DB.Database,
	}
	q := connURL.Query()
	q.Add("sslmode", cfg.DB.SSLMode)
	connURL.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), connURL.String())
	if err != nil {
		logger.Error("unable to connect to database:", sl.Err(err))
	}

	connConfig := conn.Config()

	db := stdlib.OpenDB(*connConfig)

	err = db.Ping()
	if err != nil {
		logger.Error("error pinging database:", sl.Err(err))
	}

	logger.Info("Successfully connected to PostgreSQL via database/sql!")

	return db
}
