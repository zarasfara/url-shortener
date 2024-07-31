package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/zarasfara/url-shortener/internal/logger/sl"
)

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Failed to decode request body", sl.Err(err))
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, render.M{"error": "Invalid request body"})
		return
	}

	if req.URL == "" {
		slog.Error("URL not provided")
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, render.M{"error": "URL not provided"})
		return
	}

	alias, err := h.services.UrlShortener.SaveUrl(req.URL)
	if err != nil {
		slog.Error("Failed to save URL", sl.Err(err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, render.M{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	path, err := h.services.QRCode.Save(alias, req.URL)
	if err != nil {
		slog.Error("Failed to create qr code", sl.Err(err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, render.M{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, map[string]string{"alias": alias, "path": path})
}

func (h *Handler) DisplayQRCode(w http.ResponseWriter, r *http.Request) {
	alias := chi.URLParam(r, "alias")

	path, err := h.services.QRCode.Get(alias)
	if err != nil {
		slog.Error("Failed to get QR code", sl.Err(err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, render.M{"error": "QR code not found"})
		return
	}

	file, err := os.Open(path)
	if err != nil {
		slog.Error("Failed to open QR code file", sl.Err(err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, render.M{"error": "Failed to open QR code file"})
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, path)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "short-url")

	fullUrl, err := h.services.UrlShortener.GetUrl(shortUrl)
	if err != nil {
		slog.Error("Failed to get URL", sl.Err(err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, render.M{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	http.Redirect(w, r, fullUrl, http.StatusFound)
}
