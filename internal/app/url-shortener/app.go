package url_shortener

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zarasfara/url-shortener/internal/config"
	"github.com/zarasfara/url-shortener/internal/database/postgres"
	"github.com/zarasfara/url-shortener/internal/handlers"
	"github.com/zarasfara/url-shortener/internal/logger/sl"
	"github.com/zarasfara/url-shortener/internal/repository"
	"github.com/zarasfara/url-shortener/internal/server"
	"github.com/zarasfara/url-shortener/internal/service"
)

func Run(env string) {
	cfg := config.MustLoad(env)
	slog.Info("starting url-shortener", slog.String("env", env))
	slog.Debug("debug messages are enabled")

	// Init database: postgres
	db := postgres.New(*cfg)

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
		slog.Info("server is starting", slog.String("address", fmt.Sprintf("%s:%s", cfg.HTTP.Address, cfg.HTTP.Port)))

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed to start", sl.WithError(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("shutting down server...")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", sl.WithError(err))
	}

	slog.Info("server exiting")
}
