package controllers

import "github.com/jdkruzr/web/views"
import "net/http"
import "fmt"

// Get me the "make a new user" page!
// GET /signup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gtpl"),
	}
}

// Now take that info and make a new user!
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Email is ", form.Email)
	fmt.Fprintln(w, "Password is ", form.Password)
}


// This lets gorilla/schema work

type SignupForm struct {
	Email	string `schema:"email"`
	Password string `schema:"password"`
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}
