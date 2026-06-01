package main

import (
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, ok := cfg.pages[normalizedURL]
	if ok {
		return false
	}

	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}
	return true
}

func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	cfg.pages[normalizedURL] = data
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func configure(rawBaseURL string, maxConcurrency int, maxPage int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return &config{}, err
	}

	return &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPage,
	}, nil
}
