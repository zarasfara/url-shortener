package url_shortener

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zarasfara/url-shortener/internal/config"
	"github.com/zarasfara/url-shortener/internal/database/postgres"
	"github.com/zarasfara/url-shortener/internal/handlers"
	"github.com/zarasfara/url-shortener/internal/repository"
	"github.com/zarasfara/url-shortener/internal/server"
	"github.com/zarasfara/url-shortener/internal/service"
)

func Run(env string, logger *slog.Logger) {
	cfg := config.MustLoad(env)
	logger.Info("starting url-shortener", slog.String("env", env))
	logger.Debug("debug messages are enabled")

	// Init database: postgres
	db := postgres.New(*cfg, logger)

	// Init repositories
	repos := repository.NewRepository(db)

	// Init services
	services := service.NewServices(repos)

	// Init handlers
	handler := handlers.NewHandler(services)

	// Init router: chi
	router := server.NewRouter(handler)

	// Init server
	srv := server.NewServer(cfg, router)

	go func() {
		logger.Info("server is starting", slog.String("address", cfg.HTTP.Address))

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server failed to start", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("server forced to shutdown", err)
	}

	logger.Info("server exiting")
}
