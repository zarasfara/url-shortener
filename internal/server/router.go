package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zarasfara/url-shortener/internal/handlers"
	"net/http"
)

func NewRouter(handler *handlers.Handler) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Public routes
	r.Get("/", handler.HelloWorld)
	r.Get("/qrcode/{alias}", handler.DisplayQRCode)
	r.Get("/qr/{short-url}", handler.GetQRCode)

	r.Get("/{short-url}", handler.Redirect)

	r.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Post("/shorten", handler.ShortenURL)
		})
	})

	return r
}
