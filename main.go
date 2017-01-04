package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/templates"
)

var homeTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeTemplate.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href = \"mailto:jarett@reticulum.us\">jarett@reticulum.us</a>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "These questions are FREQUENTLY ASKED!")
}

#func Handle404(w http.ResponseWriter, r *http.Request) {
#	w.Header().Set("Content-Type", "text/html")
#	fmt.Fprint(w, "Nope. 404.")
#}

func main() {

	var err error
	homeTemplate, err = template.ParseFiles("views/home.gtpl")
	if err != nil {
		panic(err)
	}
	
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("contact", contact)
	r.HandleFunc("faq", faq)
#	r.NotFoundHandler = Handle404
	http.ListenAndServe(":3000", r)
}
