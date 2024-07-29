package handlers

import (
	"encoding/json"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, render.M{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	alias, err := h.services.UrlShortener.SaveUrl(req.URL)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, render.M{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, render.M{"alias": alias})
}
