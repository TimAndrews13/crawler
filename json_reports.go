package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {

	keys := make([]string, 0, len(pages))
	for p := range pages {
		keys = append(keys, p)
	}
	sort.Strings(keys)

	sortedPages := make([]PageData, 0, len(pages))
	for _, k := range keys {
		sortedPages = append(sortedPages, pages[k])
	}

	jsonData, err := json.MarshalIndent(sortedPages, "", " ")
	if err != nil {
		fmt.Printf("error marshalling json: %v\n", err)
		return err
	}

	os.WriteFile(filename, jsonData, 0644)

	return nil
}
