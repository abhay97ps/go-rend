package main

import (
    "fyne.io/fyne/v2/app"
    "log"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Simple Rendering Engine")

    url := "https://austinhenley.com/blog.html" // Replace with the URL you want to fetch
    htmlContent, err := fetchHTML(url)
    if err != nil {
        log.Fatal(err)
    }

    parsedContent, err := parseHTML(htmlContent)
    if err != nil {
        log.Fatal(err)
    }

    // Call rendering function with parsed content
    renderContent(parsedContent, myWindow)
}
