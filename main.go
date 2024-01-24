package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseAndPrintText(htmlContent string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		panic(err)
	}

	doc.Find("body").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		fmt.Println(text)
	})
}

func fetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	url := "http://example.com" // Test URL
	htmlContent, err := fetchHTML(url)
	if err != nil {
		panic(err)
	}
	// fmt.Println(htmlContent)
	parseAndPrintText(htmlContent)
}
