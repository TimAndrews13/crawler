package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) != 4 {
		fmt.Println("3 arguments required: website, MaxConcurrency, MaxPages")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v\n", os.Args[1])

	parsedURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("error parsing base URL")
		return
	}

	maxConcurrency, _ := strconv.Atoi(os.Args[2])
	maxPages, _ := strconv.Atoi(os.Args[3])

	cfg := &config{
		pages:              make(map[string]PageData),
		baseURL:            parsedURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.wg.Add(1)

	go cfg.crawlPage(os.Args[1])

	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}

	writeJSONReport(cfg.pages, "report.json")
}
