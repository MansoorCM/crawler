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
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		return
	}

	rawURL := args[1]

	maxConcurrency := getMaxConcurrency(args, 2)
	maxPages := getMaxPages(args, 3)

	fmt.Println("starting crawl of: ", rawURL)

	cfg, err := configure(rawURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawURL)
	cfg.wg.Wait()

	fmt.Println()
	for page, count := range cfg.pages {
		fmt.Println(page, count)
	}

}
