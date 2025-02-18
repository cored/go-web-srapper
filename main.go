package main

import (
	"net/http"
	"time"
	"io"
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

func (s *Scrapper) Scrape(url string) (string, error) {
	resp, err := s.client.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
