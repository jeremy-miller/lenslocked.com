package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeremy-miller/lenslocked.com/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func main() {
	homeView = views.New("bootstrap", "views/home.gohtml")
	contactView = views.New("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "<h1>Sorry, but we couldn't find the page you were looking for</h1>")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil); err != nil {
		panic(err)
	}
}
