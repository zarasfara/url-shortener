package handlers

import (
	"github.com/go-chi/render"
	"github.com/zarasfara/url-shortener/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Services
}

func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, render.M{"message": "hello, world"})
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}
