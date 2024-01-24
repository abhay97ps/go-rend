package main

import "testing"

func TestParseHTML(t *testing.T) {
    htmlContent := `<html><body><p>Test paragraph.</p><img src="test.jpg"/></body></html>`
    parsedContent, err := parseHTML(htmlContent)
    if err != nil {
        t.Errorf("Failed to parse HTML: %v", err)
    }
    if parsedContent.Text != "Test paragraph." {
        t.Errorf("Expected to find text 'Test paragraph.', found '%v'", parsedContent.Text)
    }
    if len(parsedContent.ImageUrls) != 1 || parsedContent.ImageUrls[0] != "test.jpg" {
        t.Errorf("Expected to find image url 'test.jpg', found %v", parsedContent.ImageUrls)
    }
}
