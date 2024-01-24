package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParsedContent struct {
	Text      string
	ImageUrls []string
}

func fetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func parseHTML(htmlContent string) (*ParsedContent, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var imageUrls []string
	var textContent string

	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		textContent += s.Text()
	})

	doc.Find("img").Each(func(index int, item *goquery.Selection) {
		if imgSrc, exists := item.Attr("src"); exists {
			imageUrls = append(imageUrls, imgSrc)
		}
	})

	return &ParsedContent{
		Text:      textContent,
		ImageUrls: imageUrls,
	}, nil
}
