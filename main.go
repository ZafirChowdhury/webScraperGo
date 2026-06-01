package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("no website provided")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	if len(os.Args) > 4 {
		log.Println("too many arguments provided")
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	baseURL := os.Args[1]

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println(err.Error())
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		return
	}

	maxPage, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Println(err.Error())
		fmt.Println("usage: ./crawler URL maxConcurrency maxPages")
		return
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPage)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Printf("starting crawl of: %s...\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for normalizedURL := range cfg.pages {
		fmt.Printf("found: %s\n", normalizedURL)
	}

	err = writeJSONReport(cfg.pages, "report.json")
	if err != nil {
		log.Println(err.Error())
	}
}
