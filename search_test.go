package animetosho_test

import (
	"testing"

	"github.com/c032/animetosho-go"
)

func TestSearch(t *testing.T) {
	results, err := animetosho.Search("shakugan no shana")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("no results")
	}

	for i, result := range results {
		if result.Title == "" {
			t.Errorf("results[%d].Title = %#v; want non-empty string", i, result.Title)
		}
		if result.URL == "" {
			t.Errorf("results[%d].URL = %#v; want non-empty string", i, result.URL)
		}
		if result.MagnetURL == "" {
			t.Errorf("results[%d].MagnetURL = %#v; want non-empty string", i, result.MagnetURL)
		}
		if result.TorrentURL == "" {
			t.Errorf("results[%d].TorrentURL = %#v; want non-empty string", i, result.TorrentURL)
		}
	}
}
