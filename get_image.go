package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println("error reading html")
		return []string{}, nil
	}

	var imgUrls []string
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			parsedSrc, err := url.Parse(src)
			if err != nil {
				fmt.Println("error parsing src")
				return
			}

			resolvedSrc := baseURL.ResolveReference(parsedSrc)

			imgUrls = append(imgUrls, resolvedSrc.String())
		}
	})

	return imgUrls, nil
}
