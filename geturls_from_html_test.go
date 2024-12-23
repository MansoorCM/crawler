package main

import (
	"reflect"
	"testing"
)

func TestGetURLSFromHtml(t *testing.T) {

	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
    		<body>
        		<a href="https://blog.boot.dev"><span>Go to Boot.dev, you React Andy</span></a>
    		</body>
		</html>`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:     "relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/five">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/five"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := getURLSfromHTML(tt.inputBody, tt.inputURL)
			if err != nil {
				t.Errorf("failed to get URLs %v", err)
			} else if len(res) != len(tt.expected) {
				t.Errorf("Didn't get the expected URLs count, expected : %d, actual : %d", len(tt.expected), len(res))
			}
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("the expected URLs and actual ones doesn't match")
			}
		})
	}
}
