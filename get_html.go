package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("network error %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("got http error %w", err)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType == "" || !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("response content type %s is not text/html", contentType)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response body")
	}

	return string(body), nil
}
