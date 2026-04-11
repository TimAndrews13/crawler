package main

import (
	"testing"
)

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputHTML string
		expected  string
	}{
		{
			name:      "pull h1",
			inputHTML: "<html><body><h1>Welcome to Boot.dev</h1><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "Welcome to Boot.dev",
		},
		{
			name:      "pull h2",
			inputHTML: "<html><body><h2>Boot.dev is awesome!</h2><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "Boot.dev is awesome!",
		},
		{
			name:      "return empty string",
			inputHTML: "<html><body><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "",
		},
		{
			name:      "pull h1 when h1 and h2 are present",
			inputHTML: "<html><body><h1>Welcome to Boot.dev</h1><h2>Boot.dev is awesome!</h2><main><p>Learn to code by building real projects.</p><p>This is the second paragraph.</p></main></body></html>",
			expected:  "Welcome to Boot.dev",
		},
		// more test cases go here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected url: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
