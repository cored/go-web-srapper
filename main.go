package main

import (
	"net/http"
	"time"
)

type Scrapper struct {
	client *http.Client
}

func NewScrapper() *Scrapper {
	return &Scrapper{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
