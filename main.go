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

	fmt.Println("starting crawl of: ", args[1])
}
