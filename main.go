package main

import (
	// "fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jdkruzr/web/views"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	faqView.Template.ExecuteTemplate(w, faqView.Layout, nil)
}

func main() {

	// var err error
	
	homeView = views.NewView("bootstrap", "views/home.gtpl")
	
	contactView = views.NewView("bootstrap", "views/contact.gtpl")
	
	faqView = views.NewView("bootstrap", "views/faq.gtpl")
	
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
