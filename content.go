package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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
