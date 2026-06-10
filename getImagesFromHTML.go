package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(html string, baseURL *url.URL) ([]string, error) {
	imgURLs := []string{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("coundn't parse html: %w", err)
	}

	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		parsedSrc, err := url.Parse(src)
		if err != nil {
			return
		}

		imgURL := baseURL.ResolveReference(parsedSrc)
		imgURLs = append(imgURLs, imgURL.String())
	})

	return imgURLs, nil
}
