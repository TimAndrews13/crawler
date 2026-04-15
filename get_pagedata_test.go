package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
    </body></html>`

	actual := extractPageData(inputBody, inputURL)

	expected := PageData{
		URL:            "https://crawler-test.com",
		Heading:        "Test Title",
		FirstParagraph: "This is the first paragraph.",
		OutgoingLinks:  []string{"https://crawler-test.com/link1"},
		ImageURLs:      []string{"https://crawler-test.com/image1.jpg"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestExtractPageDataOne(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Test Page 1</title>
</head>
<body>
    <h1>Welcome to Test Page 1</h1>

    <p>This is a sample paragraph for testing HTML parsing.</p>

    <img src="/images/test1-img1.png" alt="Test Image 1">
    <img src="/assets/test1-img2.jpg" alt="Test Image 2">

    <a href="/about">About Us</a>
    <a href="/contact">Contact</a>
</body>
</html>`

	actual := extractPageData(inputBody, inputURL)

	expected := PageData{
		URL:            "https://crawler-test.com",
		Heading:        "Welcome to Test Page 1",
		FirstParagraph: "This is a sample paragraph for testing HTML parsing.",
		OutgoingLinks: []string{
			"https://crawler-test.com/about",
			"https://crawler-test.com/contact",
		},
		ImageURLs: []string{
			"https://crawler-test.com/images/test1-img1.png",
			"https://crawler-test.com/assets/test1-img2.jpg",
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestExtractPageDataTwo(t *testing.T) {
	inputURL := "https://crawler-test.com"
	inputBody := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Test Page 2</title>
</head>
<body>
    <h2>Test Page 2 Header</h2>

    <p>This page is used to validate link and image extraction logic.</p>

    <img src="/static/img/test2-img1.png" alt="Sample Image A">
    <img src="/static/img/test2-img2.png" alt="Sample Image B">

    <a href="/home">Home</a>
    <a href="/blog">Blog</a>
</body>
</html>`

	actual := extractPageData(inputBody, inputURL)

	expected := PageData{
		URL:            "https://crawler-test.com",
		Heading:        "Test Page 2 Header",
		FirstParagraph: "This page is used to validate link and image extraction logic.",
		OutgoingLinks: []string{
			"https://crawler-test.com/home",
			"https://crawler-test.com/blog",
		},
		ImageURLs: []string{
			"https://crawler-test.com/static/img/test2-img1.png",
			"https://crawler-test.com/static/img/test2-img2.png",
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
