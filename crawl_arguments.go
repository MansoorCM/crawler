package main

import (
	"fmt"
	"strconv"
)

func getMaxConcurrency(args []string, idx int) int {

	maxConcurrency := 5

	if len(args) > idx {
		maxConcurrencyInput, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("invalid maxConcurrency value, setting it to default")
			fmt.Println()
		} else {
			maxConcurrency = maxConcurrencyInput
		}
	}

	return maxConcurrency
}

func getMaxPages(args []string, idx int) int {

	maxPages := 100

	if len(args) > idx {
		maxPagesInput, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("invalid maxPages value, setting it to default")
			fmt.Println()
		} else {
			maxPages = maxPagesInput
		}
	}

	return maxPages
}
