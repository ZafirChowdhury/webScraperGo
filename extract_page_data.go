package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
}

func (p PageData) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "┌─ PageData ───────────────────────────────────\n")
	fmt.Fprintf(&sb, "│ URL:       %s\n", p.URL)
	fmt.Fprintf(&sb, "│ Heading:   %s\n", p.Heading)
	fmt.Fprintf(&sb, "│ Paragraph: %s\n", p.FirstParagraph)

	fmt.Fprintf(&sb, "│ Links (%d):\n", len(p.OutgoingLinks))
	for _, link := range p.OutgoingLinks {
		fmt.Fprintf(&sb, "│   → %s\n", link)
	}

	fmt.Fprintf(&sb, "│ Images (%d):\n", len(p.ImageURLs))
	for _, img := range p.ImageURLs {
		fmt.Fprintf(&sb, "│   🖼  %s\n", img)
	}

	fmt.Fprintf(&sb, "└──────────────────────────────────────────────\n")
	return sb.String()
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
