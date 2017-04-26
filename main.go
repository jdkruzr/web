package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jdkruzr/web/views"
	"github.com/jdkruzr/web/controllers"
)

const (
host = "localhost"
port = 5432
user = "web"
password = "supersecurepassword"
dbname = "web_dev"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println(homeView.Layout)
	homeView.Render(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	contactView.Render(w, nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	faqView.Render(w, nil)
}

func main() {

	staticC := controllers.NewStatic()
    usersC := controllers.NewUsers()
    galleriesC := controllers.NewGalleries()
	
	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET") // we are not calling the "New" method here, we're passing the method as an argument to the r.HandleFunc() call.
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/galleries/new", galleriesC.New).Methods("GET")
	http.ListenAndServe(":3000", r)
}
