package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Simple Rendering Engine")

	// URL entry field
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter URL...")

	// Container for content
	content := container.NewVBox()

	// Button to fetch and render content
	loadButton := widget.NewButton("Load", func() {
		url := urlEntry.Text
		renderContent(url, content)
	})

	// Add entry field and button to the window
	myWindow.SetContent(container.NewVBox(urlEntry, loadButton, content))
	myWindow.ShowAndRun()
}
