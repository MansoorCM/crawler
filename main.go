package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Web Crawler!")
	if len(args) < 2 {
		fmt.Println("no website provided")
		return
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		return
	}

	rawURL := args[1]
	fmt.Println("starting crawl of: ", rawURL)

	pages, err := crawlPage(rawURL, rawURL, make(map[string]int))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	for page, count := range pages {
		fmt.Println(page, count)
	}

}
