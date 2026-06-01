# webScraperGo

## What is this?

A concurrent web crawler written in Go. Give it a starting URL and it will crawl every page on that domain, extract data from each one, and save the results to a JSON report.

## Functionality

For each page visited, the crawler extracts:

- Page heading (h1, falls back to h2)
- First paragraph (prefers content inside `<main>`)
- All outgoing links
- All image URLs

It stays within the original domain, deduplicates visited pages, and runs goroutines concurrently to speed things up.

## How to run

```bash
go build -o crawler
./crawler <URL> <maxConcurrency> <maxPages>
```

**Example:**
```bash
./crawler https://example.com 5 50
```

- `URL` — the starting page to crawl
- `maxConcurrency` — how many pages to fetch at the same time
- `maxPages` — stop after crawling this many pages

Results are saved to `report.json` in the current directory.

## What I learned building this

- **Goroutines and sync primitives** — using `sync.WaitGroup` to wait for all crawls to finish and `sync.Mutex` to safely write to a shared map from multiple goroutines
- **Channels as semaphores** — a buffered channel (`chan struct{}`) is a clean way to cap the number of goroutines running at once
- **URL handling** — parsing, normalizing, and resolving relative URLs against a base using Go's `net/url` package
- **HTML parsing with goquery** — selecting and extracting content from HTML using CSS-style selectors
