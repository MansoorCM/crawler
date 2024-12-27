package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrURL string, pages map[string]int) (map[string]int, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return pages, fmt.Errorf("error parsing baseURL, %w", err)
	}

	currURL, err := url.Parse(rawCurrURL)
	if err != nil {
		return pages, fmt.Errorf("error parsing rawCurrURL, %w", err)
	}

	if baseURL.Hostname() != currURL.Hostname() {
		return pages, nil
	}

	normalizedURL, err := normalizeURL(rawCurrURL)
	if err != nil {
		return pages, err
	}

	if _, visited := pages[normalizedURL]; visited {
		pages[normalizedURL] += 1
		return pages, nil
	}

	pages[normalizedURL] = 1

	html, err := getHTML(rawCurrURL)
	if err != nil {
		return pages, err
	}

	fmt.Printf("crawling %s\n", normalizedURL)

	links, err := getURLSfromHTML(html, rawBaseURL)
	if err != nil {
		return pages, err
	}

	for _, link := range links {
		pages, _ = crawlPage(rawBaseURL, link, pages)
	}

	return pages, nil
}
