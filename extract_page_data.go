package main

import (
	"log"
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
	u, err := url.Parse(pageURL)
	if err != nil {
		log.Println(err.Error())
		return PageData{}
	}

	heading := getHeadingFromHTML(html)
	firstPeragraph := getFirstParagraphFromHTML(html)

	outgoingLinks, err := getURLsFromHTML(html, u)
	if err != nil {
		log.Println(err.Error())
		return PageData{}
	}

	imageUrls, err := getImagesFromHTML(html, u)
	if err != nil {
		log.Println(err.Error())
		return PageData{}
	}

	return PageData{
		URL:            pageURL,
		Heading:        heading,
		FirstParagraph: firstPeragraph,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageUrls,
	}
}
