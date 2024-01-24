package main

import (
	"fmt"
	"log"
)

func main() {
	url := "https://example.com/" // Replace with the URL you want to fetch
	htmlContent, err := fetchHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	parsedContent, err := parseHTML(htmlContent)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Text Content:", parsedContent.Text)
    fmt.Println("Found image URLs:", parsedContent.ImageUrls)

    // Call rendering function with parsed content
    renderContent(parsedContent)
}
