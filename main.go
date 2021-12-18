package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	// "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	// tplPath := filepath.Join("templates", "contact.gohtml")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "error parsing template", http.StatusInternalServerError)
		return
		// panic(err) //TODO remove panic
	}
	err = tpl.Execute(w, nil)
	// err = tpl.Execute(w, "test string")
	if err != nil {
		// panic(err) // TODO remove panic
		log.Printf("executing template : %v:", err)
		http.Error(w, "there was an error during execution template.", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
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

func getUser(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "userID")
	fmt.Fprint(w, "inside User using Chi ", user)
}
func getUser2(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "userID")
	fmt.Fprint(w, "inside User using Chi ", user)
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
	// var router Router
	// http.ListenAndServe
	// var router chi.NewRouter()
	r := chi.NewRouter()
	// r.Get("/user", userHandler)
	r.Use(middleware.Logger)
	r.Get("/home", homeHandler)
	r.Get("/contact", contactHandler)
	r.Route("/user", func(r chi.Router) {
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", getUser)
		})
		// r.Get("/user2", getUser2)

	})
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server at port 4500")
	err := http.ListenAndServe(":4500", r)
	if err != nil {
		panic(err)
	}
}
