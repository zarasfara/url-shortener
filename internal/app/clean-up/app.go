package cleanup

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zarasfara/url-shortener/internal/config"
	"github.com/zarasfara/url-shortener/internal/database/postgres"
	"github.com/zarasfara/url-shortener/internal/qrcode"
)

const monthsOld = 1

func DeleteOldRecords() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	cfg := config.MustLoad(env)

	db := postgres.New(*cfg)
	defer db.Close()

	cutoffDate := time.Now().AddDate(0, -monthsOld, 0)

	_, err = db.Exec("DELETE FROM links WHERE created_at < $1", cutoffDate)
	if err != nil {
		log.Printf("error deleting old records: %v", err)
	} else {
		log.Println("old records successfully deleted")
	}

	err = qrcode.DeleteOldQRCodeFiles(cutoffDate)
	if err != nil {
		log.Printf("error deleting old QR codes: %v", err)
	} else {
		log.Println("old QR codes successfully deleted")
	}
}

