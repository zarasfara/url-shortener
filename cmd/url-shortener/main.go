package main

import (
	"log"
	"os"

	_ "github.com/zarasfara/url-shortener/docs"
	urlshortener "github.com/zarasfara/url-shortener/internal/app/url-shortener"
	"github.com/zarasfara/url-shortener/internal/logger"

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

	logger.NewLogger(env)

	urlshortener.Run(env)
}