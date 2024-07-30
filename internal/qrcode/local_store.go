package qrcode

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/skip2/go-qrcode"
)

const uploadDirectory = "uploads"

// LocalStore is an implementation of the Store interface that saves QR codes locally.
type LocalStore struct {
}

// NewLocalStore creates a new LocalStore with the given directory.
func NewLocalStore() *LocalStore {
	return &LocalStore{}
}

// Save saves the QR code data to a file and returns the file path.
func (s *LocalStore) Save(alias, urlToSave string) (string, error) {
	// Check if the directory exists
	if _, err := os.Stat(uploadDirectory); os.IsNotExist(err) {
		// Create the directory
		if err := os.MkdirAll(uploadDirectory, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create directory %s: %v", uploadDirectory, err)
		}
	}

	// Generate the path to the QR code file
	qrPath := filepath.Join(uploadDirectory, fmt.Sprintf("%s.png", alias))

	// Generate and save the QR code
	if err := qrcode.WriteFile(urlToSave, qrcode.Medium, 256, qrPath); err != nil {
		return "", err
	}

	return qrPath, nil
}

// Get returns the file path of the QR code for the given alias.
func (s *LocalStore) Get(alias string) (string, error) {
	qrPath := filepath.Join(uploadDirectory, fmt.Sprintf("%s.png", alias))
	if _, err := os.Stat(qrPath); os.IsNotExist(err) {
		return "", fmt.Errorf("QR code with alias %s not found", alias)
	}
	return qrPath, nil
}

// Delete removes the QR code file for the given alias.
func (s *LocalStore) Delete(alias string) error {
	qrPath := filepath.Join(uploadDirectory, fmt.Sprintf("%s.png", alias))
	if _, err := os.Stat(qrPath); os.IsNotExist(err) {
		return fmt.Errorf("QR code with alias %s not found", alias)
	}
	if err := os.Remove(qrPath); err != nil {
		return fmt.Errorf("failed to delete QR code with alias %s: %v", alias, err)
	}
	return nil
}
