package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{name: "Remove scheme https", inputURL: "https://go.dev/doc", expected: "go.dev/doc"},
		{name: "Remove scheme http", inputURL: "http://go.dev/doc", expected: "go.dev/doc"},
		{name: "Remove trailing slash https", inputURL: "https://go.dev/doc/", expected: "go.dev/doc"},
		{name: "Remove trailing slash http", inputURL: "http://go.dev/doc/", expected: "go.dev/doc"},
		{name: "convert to Lower Case", inputURL: "http://Go.Dev/doc/", expected: "go.dev/doc"},
		{name: "invalid URL", inputURL: ":\\an invalid URL", expected: "", errorContains: "error parsing URL"},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := normalizeURL(tt.inputURL)
			if err != nil && !strings.Contains(err.Error(), tt.errorContains) {
				t.Errorf("Test %d - %s, FAIL : unexpected error: %v", i, tt.name, err)
			} else if err != nil && tt.errorContains == "" {
				t.Errorf("Test %d - %s, FAIL : unexpected error: %v", i, tt.name, err)
			} else if err == nil && tt.errorContains != "" {
				t.Errorf("Test %d - %s, FAIL : expected error: %s, got nothing", i, tt.name, tt.errorContains)
			}

			if res != tt.expected {
				t.Errorf("Test %d - %s, FAIL : expected URL: %s, actual: %s", i, tt.name, tt.expected, res)
			}
		})
	}
}
