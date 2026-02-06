package main

import (
	"fmt"
	"net/http"
)

var htmlContent = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>%s</title>
</head>
<body>
	<h1>%s</h1>
	<p>%s</p>
</body>
</html>
`

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	title := "Welcome to the Home Page"
	header := "Hello, World!"
	paragraph := "This is a simple web server built with Go."
	content := fmt.Sprintf(htmlContent, title, header, paragraph)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(content))
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	title := "About Us"
	header := "About Our Company"
	paragraph := "We are a leading company in the tech industry, providing innovative solutions to our clients."
	content := fmt.Sprintf(htmlContent, title, header, paragraph)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(content))
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	title := "Contact Us"
	header := "Get in Touch"
	paragraph := "Feel free to reach out to us via email at xyz@gmail.com"
	content := fmt.Sprintf(htmlContent, title, header, paragraph)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(content))
}
