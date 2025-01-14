package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
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

	printReport(cfg.pages, rawBaseURL)

	apiKey := os.Getenv("PERPLEXITY_KEY")
	if apiKey == "" {
		fmt.Println("invalid api key for LLM")
		return
	}
	hosts := []string{"www.chess.com", "www.google.com", "www.openai.com", "www.perplexity.ai", "news.ycombinator.com"}
	printReportFromLLM(apiKey, hosts)
}
