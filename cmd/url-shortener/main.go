package main

import (
	"log"
	"os"

	_ "github.com/zarasfara/url-shortener/docs"
	urlshortener "github.com/zarasfara/url-shortener/internal/app/url-shortener"

	"log/slog"

	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

//	@title			Swagger Example API
//	@description	An application for shortening links.
//	@version		1.0
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = envDev
	}

	logger := setupLogger(env)

	urlshortener.Run(env, logger)
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}
