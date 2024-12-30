package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {

	pages_slice := getPagesSliceFromMap(pages)
	sortPagesDescendingCount(pages_slice)
	printReportHelper(pages_slice, baseURL)

}

func getPagesSliceFromMap(pages map[string]int) []Page {
	pages_slice := make([]Page, len(pages))

	for page, count := range pages {
		pages_slice = append(pages_slice, Page{Link: page, Count: count})
	}

	return pages_slice
}

func sortPagesDescendingCount(pages_slice []Page) {
	sort.Slice(pages_slice, func(i, j int) bool {
		if pages_slice[i].Count == pages_slice[j].Count {
			return pages_slice[i].Link < pages_slice[j].Link
		}
		return pages_slice[i].Count > pages_slice[j].Count
	})
}

func printReportHelper(pages_slice []Page, baseURL string) {
	fmt.Println("=============================")
	fmt.Println("  REPORT for ", baseURL)
	fmt.Println("=============================")

	for _, page := range pages_slice {
		if page.Count == 0 {
			continue
		}
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.Link)
	}
}

type Page struct {
	Link  string
	Count int
}
