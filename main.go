package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Check your whatsapp<h1/>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("data-1", "xx")
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		pageNotFound(w, r)

	}
}

func main() {
	// fmt.Sprintf()
	// http.ListenAndServe
	// http.HandleFunc("/", homeHandler)
	http.HandleFunc("/", pathHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server at port 8000")
	http.ListenAndServe(":8000", nil)
}
