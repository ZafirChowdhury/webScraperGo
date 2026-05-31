package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	header := doc.Find("h1").First().Text()
	if header == "" {
		header = doc.Find("h2").First().Text()
	}

	return header
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	main := doc.Find("main").First()
	p := main.Find("p").First().Text()
	if p == "" {
		p = doc.Find("p").First().Text()
	}

	return p
}
