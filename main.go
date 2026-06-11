package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]PageData
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	args := os.Args
	argsWithoutProg := args[1:]

	if len(argsWithoutProg) < 3 {
		fmt.Print("very few arguments\n")
		os.Exit(1)
	}

	if len(argsWithoutProg) > 3 {
		fmt.Print("too many arguments provided\n")
		os.Exit(1)
	}

	providedURL := argsWithoutProg[0]
	providedConcurrency := argsWithoutProg[1]
	providedMaxPages := argsWithoutProg[2]

	providedConcurrencyInt, err := strconv.Atoi(providedConcurrency)
	if err != nil {
		log.Fatal("couldn't convert to integer")
	}

	providedMaxPagesInt, err := strconv.Atoi(providedMaxPages)
	if err != nil {
		log.Fatal("couldn't convert to integer")
	}

	cfg, err := configure(providedURL, providedConcurrencyInt, providedMaxPagesInt)
	if err != nil {
		log.Fatalf("couldn't configure from main, error: %v", err)
	}
	fmt.Printf("starting crawl of %s", providedURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(providedURL)
	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}

	if err := writeJSONReport(cfg.pages, "report.json"); err != nil {
		log.Fatal("couldn't write JSON report")
	}
	os.Exit(0)
}
