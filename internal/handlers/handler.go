package handlers

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/zarasfara/url-shortener/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, render.M{"message": "hello, world"})
}
