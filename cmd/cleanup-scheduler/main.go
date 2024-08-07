package main

import (
	"log"

	"github.com/robfig/cron/v3"
	cleanup "github.com/zarasfara/url-shortener/internal/app/clean-up"
)
func main() {
	c := cron.New()
	_, err := c.AddFunc("*/3 * * * *", cleanup.DeleteOldRecords)
	if err != nil {
		log.Fatalf("error adding function to cron: %v", err)
	}
	c.Start()

	log.Println("scheduler started. Waiting for termination signal...")

	select {}
}

