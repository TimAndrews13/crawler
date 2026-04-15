package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v\n", os.Args[1])
	pages := make(map[string]int)
	crawlPage(os.Args[1], os.Args[1], pages)

	fmt.Printf("pages: %v\n", pages)
}
