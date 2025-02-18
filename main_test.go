package main

import (
	"testing"
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
