package memory

import (
	"github.com/zarasfara/url-shortener/internal/repository/errors"
)


type InMemoryStorage struct {
	storage map[string]any
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		storage: make(map[string]any),
	}
}

func (ms *InMemoryStorage) SaveUrl(url, alias string) error {
	// if alias already exists
	if _, ok := ms.storage[alias]; ok {
		return errors.ErrAliasAlreadyExists
	}

	ms.storage[alias] = url

	return nil
}

func (ms *InMemoryStorage) GetURL(url string) (string, error) {
	panic("implement me")
}
