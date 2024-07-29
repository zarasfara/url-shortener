package service

import "github.com/zarasfara/url-shortener/internal/repository"

type Services struct {
	UrlShortener UrlShortenerService
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		UrlShortener: newUrlShortenerService(repos.UrlShortenerStorage),
	}
}
