package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	// "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/timereversal/lenslocked/controllers"
	"github.com/timereversal/lenslocked/templates"
	"github.com/timereversal/lenslocked/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	// tplPath := filepath.Join("templates", "contact.gohtml")
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was an error Parsing the template.", http.StatusInternalServerError)
		return
	}
	// viewTpl := views.Template{
	// 	HTMLTpl: tpl,
	// }
	// viewTpl.Execute(w, nil)
	t.Execute(w, nil)

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
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)

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
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	tpl, err := views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	// tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	tpl, err = views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	// tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	tpl, err = views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.FAQ(tpl))

	tpl, err = views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/signup", controllers.FAQ(tpl))
	// tpl, err = views.Parse(filepath.Join("templates", "about.gohtml"))
	tpl, err = views.ParseFS(templates.FS, "about.gohtml", "tailwind.gohtml")
	if err != nil {
		panic(err)
	}
	r.Get("/about", controllers.StaticHandler(tpl))

	// fmt.Sprintf()
	// var router Router
	// http.ListenAndServe
	// var router chi.NewRouter()
	// r.Get("/user", userHandler)
	// r.Get("/contact", contactHandler)
	// r.Get("/faq", faqHandler)
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
	err = http.ListenAndServe(":4500", r)
	if err != nil {
		panic(err)
	}
}
