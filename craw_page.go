package main

import (
	"fmt"
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if currURL.Hostname() != cfg.baseURL.Hostname() {
		fmt.Printf("Skiping url %s as it is in diffrent domain\n", rawCurrentURL)
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	pageData := extractPageData(htmlBody, rawCurrentURL)
	cfg.setPageData(normalizedURL, pageData)

	for _, nextURL := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
