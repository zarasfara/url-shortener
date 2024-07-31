package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/zarasfara/url-shortener/docs"
	"github.com/zarasfara/url-shortener/internal/handlers"
	"github.com/zarasfara/url-shortener/internal/utils"
)

// NewRouter initializes and returns a new router
func NewRouter(handler *handlers.Handler) http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Swagger documentation
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	// Static files for QR codes
	utils.FileServer(r, "/uploads", http.Dir("./uploads"))

	// Routes
	r.Get("/", handler.HelloWorld)
	r.Get("/{alias}", handler.Redirect)

	r.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Post("/shorten", handler.ShortenURL)
		})
	})

	return r
}
