package main

import "testing"

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
