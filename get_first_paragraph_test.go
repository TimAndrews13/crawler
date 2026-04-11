package main

import (
	"testing"
)

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "pull p1 from main",
			inputHTML: "<html><body><p>Outside paragraph.</p><main><p>Main paragraph.</p></main></body></html>",
			expected:  "Main paragraph.",
		},
		{
			name:      "pull p1 wiht no main",
			inputHTML: "<html><body><p>Outside paragraph.</p><p>Main paragraph.</p></body></html>",
			expected:  "Outside paragraph.",
		},
		{
			name:      "return empty string",
			inputHTML: "<html><body><h1>Welcome to Boot.dev</h1><h2>Boot.dev is awesome!</h2><main></main></body></html>",
			expected:  "",
		},
		// more test cases go here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected url: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
