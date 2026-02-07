package main

import (
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
	app.render(w, "index.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	app.render(w, "login.html", nil)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	app.render(w, "register.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	app.render(w, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	app.render(w, "contact.html", nil)
}
