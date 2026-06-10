package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(html string, baseURL *url.URL) ([]string, error) {
	URLs := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("unable to parse HTML: %w", err)
	}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}
		parsedHref, err := url.Parse(href)
		if err != nil {
			return
		}

		absoluteURL := baseURL.ResolveReference(parsedHref)
		URLs = append(URLs, absoluteURL.String())
	})

	return URLs, nil
}
