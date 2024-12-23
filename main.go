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
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawURL := args[1]
	fmt.Println("starting crawl of: ", rawURL)

	htmlBody, err := getHTML(rawURL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(htmlBody)

}
