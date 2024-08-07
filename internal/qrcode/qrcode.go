package qrcode

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func DeleteOldQRCodeFiles(cutoffDate time.Time) error {
	return filepath.WalkDir(UploadDirectory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fileInfo, err := os.Stat(path)
			if err != nil {
				return err
			}
			if fileInfo.ModTime().Before(cutoffDate) {
				log.Printf("Удаление старого файла QR-кода: %s", path)
				return os.Remove(path)
			}
		}
		return nil
	})
}