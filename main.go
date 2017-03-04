package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jdkruzr/web/views"
	"github.com/jdkruzr/web/controllers"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View
// var signUpView *views.View

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

/* func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(“Content-Type”, “text/html”)
	signUpView.Render(w, nil)
} */ 

func main() {

	// var err error
	
//	homeView = views.NewView("bootstrap", "views/home.gtpl")
	
//	contactView = views.NewView("bootstrap", "views/contact.gtpl")
	
//	faqView = views.NewView("bootstrap", "views/faq.gtpl")
	

	// signUpView = views.NewView("bootstrap", "views/signup.gtpl")
	staticC := controllers.NewStatic()
  usersC := controllers.NewUsers()
	
	r := mux.NewRouter()

	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET") // we are not calling the "New" method here, we're passing the method as an argument to the r.HandleFunc() call.
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
