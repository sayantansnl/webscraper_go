package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://crawler-test.com/path",
			expected: "crawler-test.com/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://crawler-test.com/path/",
			expected: "crawler-test.com/path",
		},
		{
			name:     "lowercase capital letters",
			inputURL: "https://CRAWLER-TEST.com/PATH",
			expected: "crawler-test.com/path",
		},
		{
			name:     "remove scheme and capitals and trailing slash",
			inputURL: "http://CRAWLER-TEST.com/path/",
			expected: "crawler-test.com/path",
		},
		{
			name:          "handle invalid URL",
			inputURL:      `:\\invalidURL`,
			expected:      "",
			errorContains: "couldn't parse URL",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	inputBody := "<html><body><h1>Test Title</h1></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetH1FromHTML(t *testing.T) {
	inputBody := "<html><body><h1>Test Title</h1><h2>Another Heading</h2></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetH2FromHTML(t *testing.T) {
	inputBody := "<html><body><h2>Refined land!</h2><h3>Test Title</h3></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := "Refined land!"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetEmptyString(t *testing.T) {
	inputBody := "<html><body><h3>Refined land!</h3><h4>Test Title</h4></body></html>"
	actual := getHeadingFromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestGetEmptyStringFromHTML(t *testing.T) {
	inputBody := `
		<html>
			<body>
				<h1>Billy says Oi</h1>
			</body>
		</html>
	`

	actual := getFirstParagraphFromHTML(inputBody)
	expected := ""

	if actual != expected {
		t.Errorf("expected: %q, got : %q", expected, actual)
	}
}

func TestGetParagraphFromOutsideMain(t *testing.T) {
	inputBody := `
		<html>
			<body>
				<p>You will get this</p>
				<p>You won't get this</p>
				<main>
					<h1>You won't get it</h1>
				</main>
			</body>
		</html>
	`

	actual := getFirstParagraphFromHTML(inputBody)
	expected := "You will get this"

	if actual != expected {
		t.Errorf("Expected: %q, Got: %q", expected, actual)
	}
}

func TestGetFirstParagraphFromHTMLFallback(t *testing.T) {
	inputBody := `<html><body>
		<p>First paragraph outside main.</p>
		<p>Second paragraph outside main.</p>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "First paragraph outside main."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
