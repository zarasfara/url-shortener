package service

import (
	"errors"
	"log/slog"
	"net/url"

	"github.com/teris-io/shortid"
	"github.com/zarasfara/url-shortener/internal/logger/sl"
	"github.com/zarasfara/url-shortener/internal/repository"
)

var (
	ErrInvalidURL = errors.New("invalid URL")
)

type UrlShortenerService interface {
	SaveUrl(url string) (string, error)
	GetUrl(shortUrl string) (string, error)
}

type urlShortenerService struct {
	repo repository.UrlShortenerStorage
}

func newUrlShortenerService(repo repository.UrlShortenerStorage) UrlShortenerService {
	return &urlShortenerService{
		repo: repo,
	}
}

func isValidURL(inputURL string) bool {
	_, err := url.ParseRequestURI(inputURL)
	return err == nil && inputURL != ""
}

func (s *urlShortenerService) SaveUrl(url string) (string, error) {
	if !isValidURL(url) {
		return "", ErrInvalidURL
	}

	alias, err := shortid.Generate()
	if err != nil {
		slog.Error("Failed to generate shortid", sl.Err(err))
		return "", err
	}

	err = s.repo.SaveUrl(url, alias)
	if err != nil {
		slog.Error("Failed to save URL in repository", sl.Err(err))
		return "", err
	}

	return alias, nil
}

func (s *urlShortenerService) GetUrl(shortUrl string) (string, error) {
	url, err := s.repo.GetUrl(shortUrl)
	if err != nil {
		slog.Error("Failed to retrieve URL from repository", sl.Err(err))
		return "", err
	}
	return url, nil
}
