// main.go
package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/zarasfara/url-shortener/internal/config"
	"github.com/zarasfara/url-shortener/internal/repository"
	"github.com/zarasfara/url-shortener/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = envLocal
	}

	// init logger
	log := setupLogger(env)

	// init config
	cfg := config.MustLoad(env)
	log.Info("starting url-shortener", slog.String("env", env))
	log.Debug("debug messages are enabled")

	// init storage: postgres
	db := postgres.New(*cfg, log)

	// init repositories
	_ = repository.NewRepository(db)

	// TODO: init services

	// TODO: init router: chi

	// TODO: init server
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
