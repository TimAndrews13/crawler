package main

import (
	"fmt"
	"net/url"
	"path"
)

func normalizeURL(urlString string) (string, error) {
	parts, err := url.Parse(urlString)
	if err != nil {
		fmt.Printf("error parsing url: %v", err)
		return "", err
	}
	hostName := parts.Hostname()
	path := path.Clean(parts.Path)
	if path == "." {
		path = ""
	}
	nURL := hostName + path
	return nURL, nil
}
