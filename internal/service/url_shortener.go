package service

import (
	"github.com/teris-io/shortid"
	"github.com/zarasfara/url-shortener/internal/repository"
)

type UrlShortenerService interface {
	SaveUrl(url string) (string, error)
}

type urlShortenerService struct {
	repo repository.UrlShortenerStorage
}

func newUrlShortenerService(repo repository.UrlShortenerStorage) UrlShortenerService {
	return &urlShortenerService{
		repo: repo,
	}
}

func (s *urlShortenerService) SaveUrl(url string) (string, error) {
	alias, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	err = s.repo.SaveUrl(url, alias)
	if err != nil {
		return "", err
	}

	return alias, nil
}
