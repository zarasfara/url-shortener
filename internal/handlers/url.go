package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/zarasfara/url-shortener/internal/logger/sl"
	"github.com/zarasfara/url-shortener/internal/service"
)

// ShortenURLResponse represents the response structure for ShortenURL
type ShortenURLResponse struct {
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

// ShortenURL godoc
//
//	@Summary		Shorten a URL
//	@Description	Takes a URL and returns a shortened URL alias and QR code path
//	@Tags			URL Shortener
//	@Accept			json
//	@Produce		json
//	@Param			request	body		handlers.ShortenURL.RequestBody	true	"URL to be shortened"
//	@Success		201		{object}	ShortenURLResponse
//	@Failure		400		{object}	HttpError
//	@Failure		500		{object}	HttpError
//	@Router			/api/v1/shorten [post]
func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		URL string `json:"url"`
	}

	var requestBody RequestBody

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		slog.Error("Failed to decode request body", sl.WithError(err))
		NewHttpError("Invalid request body").Write(w, http.StatusBadRequest)
		return
	}

	if requestBody.URL == "" {
		slog.Error("URL not provided")

		NewHttpError("URL not provided").Write(w, http.StatusBadRequest)
		return
	}

	alias, err := h.services.UrlShortener.SaveUrl(requestBody.URL)
	if err != nil {
		if errors.Is(err, service.ErrInvalidURL) {
			slog.Error("Invalid URL provided", sl.WithError(err))
			NewHttpError("Invalid URL").Write(w, http.StatusBadRequest)

		} else {
			slog.Error("Failed to save URL", sl.WithError(err))
			NewHttpError(http.StatusText(http.StatusInternalServerError)).Write(w, http.StatusInternalServerError)
		}
		return
	}

	path, err := h.services.QRCode.Save(alias, requestBody.URL)
	if err != nil {
		slog.Error("Failed to create QR code", sl.WithError(err))
		NewHttpError(http.StatusText(http.StatusInternalServerError)).Write(w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, ShortenURLResponse{Alias: alias, Path: path})
}

// Redirect godoc
//
//	@Summary		Redirect to the full URL
//	@Description	Redirects to the original full URL based on the shortened URL alias
//	@Tags			URL Redirect
//	@Param			alias	path	string	true	"Shortened URL alias"
//	@Success		302
//	@Failure		500	{object}	HttpError
//	@Router			/{alias} [get]
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "alias")

	fullUrl, err := h.services.UrlShortener.GetUrl(shortUrl)
	if err != nil {
		slog.Error("Failed to get URL", sl.WithError(err))
		NewHttpError(http.StatusText(http.StatusInternalServerError)).Write(w, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fullUrl, http.StatusFound)
}
