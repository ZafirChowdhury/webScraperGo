package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		log.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseURL)

	html, err := getHTML(baseURL)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(html)
}
