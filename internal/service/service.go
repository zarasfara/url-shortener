package service

import (
	"github.com/zarasfara/url-shortener/internal/qrcode"
	"github.com/zarasfara/url-shortener/internal/repository"
)

type Services struct {
	UrlShortener UrlShortenerService
	QRCode       *QRCodeService
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		UrlShortener: newUrlShortenerService(repos.UrlShortenerStorage),
		QRCode:       NewQRCodeService(qrcode.NewLocalStore()),
	}
}
