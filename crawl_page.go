package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
	}()
	defer cfg.wg.Done()
	//check maxPages
	if len(cfg.pages) >= cfg.maxPages {
		return
	}

	//make sure domain matches between rawBaseURL and rawCurrentURL
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing current url")
		return
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	//normalize current url string
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error normalizing current url")
		return
	}

	//check againt pages map, if exists increment by 1, else add to map with value of 1
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	//get html from current URL and print
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing html from current url")
		return
	}

	pageData := extractPageData(html, rawCurrentURL)
	cfg.setPageData(normalizedURL, pageData)

	fmt.Printf("crawling %s\n", rawCurrentURL)

	for _, nextURL := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go func(url string) {
			cfg.crawlPage(url)
		}(nextURL)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normalizedURL]; exists {
		return false
	} else {
		cfg.pages[normalizedURL] = PageData{}
		return true
	}

}
