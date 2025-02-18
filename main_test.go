package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestNewScrapper(t *testing.T) {
	scrapper := NewScrapper()
	if scrapper == nil {
		t.Error("Expected scrapper to be created, but got nil")
	}

	if scrapper.client == nil {
		t.Error("Expected scrapper.client to be created, but got nil")
	}
}

func TestScrape(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body>Hello, world!</body></html>"))
	}))

	defer server.Close()

	scrapper := NewScrapper()
	content, err := scrapper.Scrape(server.URL)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !strings.Contains(content, "Hello, world!") {
		t.Errorf("Expected content to contain 'Hello, world!', but got %s", content)
	}


}

func TestScrapeConcurrent(t *testing.T) {
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body>Hello, from server 1!</body></html>"))
	}))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body>Hello, from server 2!</body></html>"))
	}))

	defer server2.Close()

	urls := []string{server1.URL, server2.URL}

	scraper := NewScrapper()
	results := scraper.ScrapeConcurrent(urls)

	if len(results) != 2 {
		t.Errorf("Expected 2 results, but got %d", len(results))
	}

}
