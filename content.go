package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	h1 := doc.Find("h1, h2").First().Text()
	return strings.TrimSpace(h1)
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	var p string

	main := doc.Find("main")
	mainP := main.Find("p")
	if mainP.Length() > 0 {
		p = mainP.First().Text()
	} else {
		p = doc.Find("p").First().Text()
	}

	return strings.TrimSpace(p)
}
