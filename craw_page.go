package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	if rawCurrentURL == "" {
		rawCurrentURL = rawBaseURL
	}

	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if currURL.Hostname() != baseURL.Hostname() {
		log.Printf("Skiping url %s as it is in diffrent domain\n", rawCurrentURL)
		return
	}

	normalURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	_, ok := pages[normalURL]
	if ok {
		pages[normalURL]++
		return
	}

	pages[normalURL] = 1

	fmt.Printf("Crawling %s\n", normalURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	pageData := extractPageData(html, rawCurrentURL)
	fmt.Println(pageData.String())

	// recurcive crawl
	urls, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, u := range urls {
		crawlPage(rawBaseURL, u, pages)
	}
}
