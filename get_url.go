package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println("error reading html")
		return []string{}, nil
	}

	var urls []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			parsedHref, err := url.Parse(href)
			if err != nil {
				fmt.Println("error parsing href")
				return
			}

			resolvedURL := baseURL.ResolveReference(parsedHref)

			urls = append(urls, resolvedURL.String())
		}
	})

	return urls, nil

}
