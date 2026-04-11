package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("error reading html")
		return ""
	}

	headerF := doc.Find("h1")
	if headerF.Text() == "" {
		headerS := doc.Find("h2")
		if headerS.Text() == "" {
			return ""
		}
		return headerS.Text()
	}
	return headerF.Text()
}
