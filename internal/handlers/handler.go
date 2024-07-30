package handlers

import (
	"fmt"
	"image/png"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/skip2/go-qrcode"
	"github.com/zarasfara/url-shortener/internal/logger/sl"
	"github.com/zarasfara/url-shortener/internal/service"
)

type Handler struct {
	services *service.Services
}

func (h *Handler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, render.M{"message": "hello, world"})
}

func GenerateQRCode(url string, filePath string) error {
	// Убедитесь, что директория существует или создайте её
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Создание QR-кода
	qr, err := qrcode.New(url, qrcode.Low)
	if err != nil {
		return err
	}

	// Открытие файла для записи
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Кодирование QR-кода в PNG
	if err := png.Encode(file, qr.Image(256)); err != nil {
		return err
	}

	return nil
}

func (h *Handler) GetQRCode(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "short-url")

	fullUrl, err := h.services.UrlShortener.GetUrl(shortUrl)
	if err != nil {
		slog.Error("Failed to get URL", sl.Err(err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, render.M{"error": "Internal server error"})
		return
	}

	qrFilePath := filepath.Join("qr_codes", shortUrl+".png")

	if err := GenerateQRCode(fullUrl, qrFilePath); err != nil {
		slog.Error("Failed to generate QR code", sl.Err(err))
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, render.M{"error": "Internal server error"})
		return
	}

	http.ServeFile(w, r, qrFilePath)
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}
