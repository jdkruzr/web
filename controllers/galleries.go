package controllers

import "github.com/jdkruzr/web/views"
import "net/http"
import "fmt"

// type SignupForm struct {
//	Email	string `schema:"email"`
//	Password string `schema:"password"`
//}

type Galleries struct {
	NewView *views.View
}

func NewGalleries() *Galleries {
	return &Galleries{
		NewView: views.NewView("bootstrap", "galleries/new"),
	}
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "zomg")
	//g.NewView.Render(w, nil)
}