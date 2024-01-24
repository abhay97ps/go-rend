package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"net/http"
)

// Helper function to fetch and return an image from a URL
func fetchImage(url string) image.Image {
	resp, err := http.Get(url)
	if err != nil {
		return image.NewUniform(color.RGBA{0, 0, 0, 0}) // Return a blank image in case of error
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return image.NewUniform(color.RGBA{0, 0, 0, 0}) // Return a blank image in case of error
	}

	return img
}

// renderContent displays text and images in a simple window
func renderContent(parsedContent *ParsedContent) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Simple Rendering Engine")

	// Container to hold both text and images
	content := container.NewVBox()

	// Adding text content
	textLabel := widget.NewLabel(parsedContent.Text)
	content.Add(textLabel)

	// Adding images
	for _, imgUrl := range parsedContent.ImageUrls {
		img := fetchImage(imgUrl)
		imageWidget := canvas.NewImageFromImage(img)
		imageWidget.FillMode = canvas.ImageFillContain // Adjust the fill mode as needed
		content.Add(imageWidget)
	}

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
