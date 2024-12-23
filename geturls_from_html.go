package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLSfromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	var links []string

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}

	links = traverse(doc, links, rawBaseURL)
	return links, nil
}

func traverse(n *html.Node, links []string, rawBaseURL string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				rawUrl := attr.Val
				if !isAbsoluteURL(rawUrl) {
					rawUrl = rawBaseURL + rawUrl
				}
				links = append(links, rawUrl)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = traverse(c, links, rawBaseURL)
	}

	return links
}

func isAbsoluteURL(rawURL string) bool {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	return parsed.Scheme != "" && parsed.Host != ""
}
