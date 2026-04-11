package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("error reading html")
		return ""
	}

	mainText := doc.Find("main")
	if mainText.Text() == "" {
		firstParagraph := doc.Find("p").First()
		if firstParagraph.Text() == "" {
			return ""
		}
		return firstParagraph.Text()
	}
	firstParagraph := mainText.Find("p").First()
	return firstParagraph.Text()
}
