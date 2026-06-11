package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Fatalf("couldn't parse baseURL: %v", err)
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Fatalf("couldn't parse currentURL: %v", err)
	}

	if baseURL.Host != currentURL.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatalf("unable to normalize current URL %s, error: %v", rawCurrentURL, err)
	}

	if val, ok := pages[normalizedURL]; ok {
		pages[normalizedURL] = val + 1
		return
	}

	pages[normalizedURL] = 1

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Fatalf("couldn't get HTML from %s due to error: %v", rawCurrentURL, err)
	}

	fmt.Printf("\nCrawling %s", rawCurrentURL)

	URLs, err := getURLsFromHTML(htmlBody, baseURL)
	if err != nil {
		log.Fatalf("unable to retrieve URLs from %s due to error: %v", rawCurrentURL, err)
	}

	for _, URL := range URLs {
		crawlPage(rawBaseURL, URL, pages)
	}
}
