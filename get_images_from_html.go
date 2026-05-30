package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var imgURL []string
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}

		src = strings.TrimSpace(src)
		if src == "" {
			return
		}

		u, err := url.Parse(src)
		if err != nil {
			log.Println(err.Error())
			return
		}

		resolved := baseURL.ResolveReference(u)
		imgURL = append(imgURL, resolved.String())
	})

	return imgURL, nil
}
