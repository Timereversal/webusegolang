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
	fmt.Fprint(w, "contactHanlder")
}
func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("data-1", "xx")
	fmt.Fprint(w, `
	<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us -
    <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
  </li>
</ul>
	`)
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
func pathHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, r.URL.Path)
	// fmt.Fprintf(w, r.URL.RawPath)
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		// pageNotFound(w, r)
		http.Error(w, "Page not found a", http.StatusNotFound)

	}
}

type Router struct {
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		// pageNotFound(w, r)
		http.Error(w, "Page not found a", http.StatusNotFound)

	}
}

func main() {
	// fmt.Sprintf()
	var router Router
	// http.ListenAndServe
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server at port 4500")
	err := http.ListenAndServe(":4500", router)
	if err != nil {
		panic(err)
	}
}
