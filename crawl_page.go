package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	//make sure domain matches between rawBaseURL and rawCurrentURL
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println("error parsing base url")
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing current url")
		return
	}
	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	//normalize current url string
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("error normalizing current url")
		return
	}

	//check againt pages map, if exists increment by 1, else add to map with value of 1
	if _, exists := pages[normalizedURL]; exists {
		pages[normalizedURL] += 1
		return
	} else {
		pages[normalizedURL] = 1
	}

	//get html from current URL and print
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println("error parsing html from current url")
		return
	}
	fmt.Printf("crawling %s\n", rawCurrentURL)

	//get urls from html
	urls, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Println("error parsing urls from current html")
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
