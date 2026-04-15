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
	fmt.Printf("crawling %s\n", rawCurrentURL)

	//get urls from html
	urls, err := getURLsFromHTML(html, cfg.baseURL)
	if err != nil {
		fmt.Println("error parsing urls from current html")
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go func(url string) {
			cfg.crawlPage(url)
		}(url)
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
