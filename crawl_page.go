package main

import (
	"fmt"
	"net/url"
)

func (cfg config) crawlPage(rawCurrURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	maxPagesVisited := cfg.alreadyVisitedMaxPages()
	if maxPagesVisited {
		return
	}

	currURL, err := url.Parse(rawCurrURL)
	if err != nil {
		fmt.Printf("error parsing rawCurrURL, %v\n", err)
		return
	}

	if cfg.baseURL.Hostname() != currURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrURL)
	if err != nil {
		fmt.Printf("error normalizing url, %v\n", err)
		return
	}

	isFirstVisit := cfg.addPageVisit(normalizedURL)

	if !isFirstVisit {
		return
	}

	html, err := getHTML(rawCurrURL)
	if err != nil {
		fmt.Printf("error getting html from URL, %v\n", err)
		return
	}

	fmt.Printf("crawling %s\n", normalizedURL)

	links, err := getURLSfromHTML(html, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("error getting URLs from html, %v\n", err)
		return
	}

	for _, link := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}

}
