package main

import (
	"image"
	"image/color"
	"math"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/draw"
)

func resizeImage(img image.Image, maxWidth, maxHeight int) image.Image {
	bounds := img.Bounds()
	hRatio := float64(maxHeight) / float64(bounds.Dy())
	wRatio := float64(maxWidth) / float64(bounds.Dx())
	ratio := math.Min(hRatio, wRatio)

	newWidth := int(float64(bounds.Dx()) * ratio)
	newHeight := int(float64(bounds.Dy()) * ratio)

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, bounds, draw.Src, nil)
	return newImg
}

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

	// Resize image if it's too large
	const maxDimension = 1000 // Example limit for width and height
	if img.Bounds().Dx() > maxDimension || img.Bounds().Dy() > maxDimension {
		img = resizeImage(img, maxDimension, maxDimension)
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
