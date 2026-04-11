package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
		},
		{
			name:     "remove https scheme and trailing /",
			inputURL: "https://www.espn.com/soccer/",
			expected: "www.espn.com/soccer",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://www.google.com",
			expected: "www.google.com",
		},
		// more test cases go here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Teest %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected url: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
