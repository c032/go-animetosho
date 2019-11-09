package animetosho

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	searchURLStr = "https://animetosho.org/search"
)

func init() {
	_, err := url.Parse(searchURLStr)
	if err != nil {
		panic(err)
	}
}

// SearchResult is a search result.
type SearchResult struct {
	Title      string
	URL        string
	MagnetURL  string
	TorrentURL string
}

// Search performs a search and returns the first page of results.
func Search(terms string) ([]SearchResult, error) {
	var (
		err error

		searchURL *url.URL
		resp      *http.Response
		doc       *goquery.Document
	)

	searchURL, err = url.Parse(searchURLStr)
	if err != nil {
		return nil, err
	}

	qs := searchURL.Query()
	qs.Set("q", terms)

	searchURL.RawQuery = qs.Encode()

	resp, err = http.Get(searchURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	doc.Find(".home_list_entry").Each(func(i int, entry *goquery.Selection) {
		var (
			err error

			infoURL    *url.URL
			magnetURL  *url.URL
			torrentURL *url.URL
		)

		titleSel := entry.Find(".link a").First()
		if titleSel.Length() != 1 {
			return
		}

		infoLink := titleSel.AttrOr("href", "")
		if infoLink == "" {
			return
		}

		infoURL, err = searchURL.Parse(infoLink)
		if err != nil {
			return
		}

		title := strings.TrimSpace(titleSel.Text())
		if title == "" {
			return
		}

		dlLink := entry.Find(".dllink").First().AttrOr("href", "")
		if dlLink == "" {
			return
		}

		torrentURL, err = searchURL.Parse(dlLink)
		if err != nil {
			return
		}

		magnetLink := entry.Find(`a[href^="magnet:"]`).First().AttrOr("href", "")
		if magnetLink == "" {
			return
		}

		magnetURL, err = searchURL.Parse(magnetLink)
		if err != nil {
			return
		}

		result := SearchResult{
			Title:      title,
			URL:        infoURL.String(),
			MagnetURL:  magnetURL.String(),
			TorrentURL: torrentURL.String(),
		}

		results = append(results, result)
	})

	return results, nil
}
