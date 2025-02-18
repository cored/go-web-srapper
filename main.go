package main

import (
	"net/http"
	"time"
	"io"
	"sync"
)

type Scrapper struct {
	client *http.Client
}

type ScrapedHTML struct {
	URL     string
	Content string
	Error   error
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

func (s *Scrapper) ScrapeConcurrent(urls []string) []ScrapedHTML {
	var wg sync.WaitGroup
	results := make([]ScrapedHTML, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			content, err := s.Scrape(url)
			results[i] = ScrapedHTML{url, content, err}
		}(i, url)
	}

	wg.Wait()

	return results
}
