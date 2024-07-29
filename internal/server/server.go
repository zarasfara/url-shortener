package server

import (
	"context"
	"github.com/zarasfara/url-shortener/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{httpServer: &http.Server{
		Addr:    ":" + cfg.HTTP.Port,
		Handler: handler,
	}}
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
