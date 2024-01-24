package main

import (
	"image"
	"image/color"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
func renderContent(parsedContent *ParsedContent, window fyne.Window) {
	content := container.NewVBox()

	// Display text
	text := widget.NewLabel(parsedContent.Text)
	content.Add(text)

	for _, imgUrl := range parsedContent.ImageUrls {
		imageWidget := canvas.NewImageFromImage(image.NewUniform(color.RGBA{0, 0, 0, 0})) // placeholder image
		content.Add(imageWidget)

		go func(url string, imgWidget *canvas.Image) {
			img := fetchImage(url)
			imgWidget.Image = img
			imgWidget.Refresh() // Refresh the image widget to show the new image
		}(imgUrl, imageWidget)
	}

	window.SetContent(content)
	window.ShowAndRun()
}
