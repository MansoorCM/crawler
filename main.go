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

	rawBaseURL := args[1]

	maxConcurrency := getMaxConcurrency(args, 2)
	maxPages := getMaxPages(args, 3)

	fmt.Println("starting crawl of: ", rawBaseURL)

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println()

	printReport(cfg.pages, rawBaseURL)

}
