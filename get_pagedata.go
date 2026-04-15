package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("error parsing pageURL")
		return PageData{}
	}

	outgoingLinks, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Println("error parsing outgoing links")
		return PageData{}
	}

	imageURLs, err := getImagesFromHTML(html, baseURL)
	if err != nil {
		fmt.Println("error parsing image urls")
		return PageData{}
	}

	pageData := PageData{
		URL:            pageURL,
		Heading:        getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
	}

	return pageData
}
