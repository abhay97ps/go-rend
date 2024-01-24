package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// renderContent displays text and images in a simple window
func renderContent(parsedContent *ParsedContent, window fyne.Window) {
	content := container.NewVBox()

	// Display text
	text := widget.NewLabel(parsedContent.Text)
	content.Add(text)

	window.SetContent(content)
	window.ShowAndRun()
}
