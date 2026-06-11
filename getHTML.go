package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("couldn't make a request: %w", err)
	}

	req.Header.Set("User-Agent", "Crawler/1.0")

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("unable to get response: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("status code not okay")
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", contentType)
	}

	htmlBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %w", err)
	}

	htmlString := string(htmlBody)

	return htmlString, nil
}
