package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"log"
)

// renderContent displays text and images in a simple window
func renderContent(url string, content *fyne.Container) {
	// Fetch and parse HTML content here (you may call fetchHTML and parseHTML functions)

	// Example:
	htmlContent, err := fetchHTML(url)
	if err != nil {
	  log.Println("Error fetching URL:", err)
	  return
	}
	
	parsedContent, err := parseHTML(htmlContent)
	if err != nil {
	  log.Println("Error parsing HTML:", err)
	  return
	}

	// Display text (replace this with actual content)
	text := widget.NewLabel(parsedContent.Text)
	content.Objects = []fyne.CanvasObject{text}
	content.Refresh()
}
