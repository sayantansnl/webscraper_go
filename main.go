package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	argsWithoutProg := args[1:]

	pages := map[string]int{}

	if len(argsWithoutProg) < 1 {
		fmt.Print("no website provided\n")
		os.Exit(1)
	}

	if len(argsWithoutProg) > 1 {
		fmt.Print("too many arguments provided\n")
		os.Exit(1)
	}

	providedURL := argsWithoutProg[0]
	fmt.Printf("starting crawl of %s", providedURL)

	crawlPage(providedURL, providedURL, pages)
	fmt.Println(pages)
	os.Exit(0)
}
