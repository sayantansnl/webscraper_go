package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://crawler-test.com"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetMultipleURLsFromHTML(t *testing.T) {
	inputURL := "https://somewebsite.com"
	inputBody := `
		<html>
			<body>
				<a href="/path/abc">
					<span>Boot.dev</span>
				</a>
				<a href="/path/xyz"></a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://somewebsite.com/path/abc",
		"https://somewebsite.com/path/xyz",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetOneRelativeOneAbsoluteURLsFromHTML(t *testing.T) {
	inputURL := "https://somewebsite.com"
	inputBody := `
		<html>
			<body>
				<a href="https://somewebsite.com">
					<span>Boot.dev</span>
				</a>
				<a href="/path/xyz"></a>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{
		"https://somewebsite.com",
		"https://somewebsite.com/path/xyz",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetEmptySliceWithoutURLFromHTML(t *testing.T) {
	inputURL := "https://somewebsite.com"
	inputBody := `
		<html>
			<body>
			</body>
		</html>
	`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(actual) != 0 {
		t.Errorf("Expected actual to be empty, got: %q", actual)
	}
}
