package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// routing - mux
// routing -> hadlers -> controllers -> handler
// GET / - homePage
// POST /users -  CreateUser

type DefaultHandler struct{}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Printf("Received request for path: %s\n", path)
	data := map[string]interface{}{
		"user":     "John Doe",
		"age":      30,
		"height":   1.75,
		"location": "New York",
	}
	bs, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bs)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {

	mux := &DefaultHandler{}

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
