package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	if len(pages) == 0 {
		log.Println("No pages to save")
		return nil
	}

	keys := make([]string, 0, len(pages))
	for key, _ := range pages {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sorted := make([]PageData, 0, len(pages))
	for _, key := range keys {
		sorted = append(sorted, pages[key])
	}

	data, err := json.MarshalIndent(sorted, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("report saved to %s...\n", filename)
	return nil
}
