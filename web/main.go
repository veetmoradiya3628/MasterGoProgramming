package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
