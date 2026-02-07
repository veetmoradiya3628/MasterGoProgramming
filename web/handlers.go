package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	// app.infoLog.Printf("Session data: %s", app.session.GetString(r, "userID"))
	app.render(w, r, "index.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	// app.session.Put(r, "userID", "joseph")

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		form := NewForm(r.PostForm)
		form.Required("email", "password").
			MaxLength("email", 255).
			MaxLength("password", 255).
			MinLength("email", 3).
			Matches("email", EmailRX)

		if !form.Valid() {
			app.errorLog.Printf("Invalid form: %+v", form.Errors)
			form.Errors.Add("generic", "The data you submitted was not valid")
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")
		app.infoLog.Printf("Logged in with email %s, password %s", email, password)
	}

	app.render(w, r, "login.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	app.render(w, r, "register.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	app.render(w, r, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	app.render(w, r, "contact.html", nil)
}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("Received request for %s", r.URL.Path)
	app.render(w, r, "submit.html", nil)
}
