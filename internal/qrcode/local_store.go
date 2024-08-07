package qrcode

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/skip2/go-qrcode"
)

const (
	UploadDirectory = "uploads"

	qrCodeExtension = ".png"
	qrCodeSize      = 256
)

// LocalStore is an implementation of the Store interface that saves QR codes locally.
type LocalStore struct{}

// NewLocalStore creates a new LocalStore with the given directory.
func NewLocalStore() *LocalStore {
	return &LocalStore{}
}

// Save saves the QR code data to a file and returns the file path.
func (s *LocalStore) Save(alias, urlToSave string) (string, error) {
	// Ensure the upload directory exists
	if _, err := os.Stat(UploadDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(UploadDirectory, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create directory %s: %v", UploadDirectory, err)
		}
	}

	// Generate the path to the QR code file
	qrPath := filepath.Join(UploadDirectory, fmt.Sprintf("%s%s", alias, qrCodeExtension))

	// Generate and save the QR code
	if err := qrcode.WriteFile(urlToSave, qrcode.Medium, qrCodeSize, qrPath); err != nil {
		return "", err
	}

	return qrPath, nil
}

// Get returns the file path of the QR code for the given alias.
func (s *LocalStore) Get(alias string) (string, error) {
	qrPath := filepath.Join(UploadDirectory, fmt.Sprintf("%s%s", alias, qrCodeExtension))
	if _, err := os.Stat(qrPath); os.IsNotExist(err) {
		return "", fmt.Errorf("QR code with alias %s not found", alias)
	}
	return qrPath, nil
}

// Delete removes the QR code file for the given alias.
func (s *LocalStore) Delete(alias string) error {
	qrPath := filepath.Join(UploadDirectory, fmt.Sprintf("%s%s", alias, qrCodeExtension))
	if _, err := os.Stat(qrPath); os.IsNotExist(err) {
		return fmt.Errorf("QR code with alias %s not found", alias)
	}
	if err := os.Remove(qrPath); err != nil {
		return fmt.Errorf("failed to delete QR code with alias %s: %v", alias, err)
	}
	return nil
}
