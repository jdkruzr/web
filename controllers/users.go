package controllers

func NewUsers() *Users {
	return &Users(
		NewView: views.NewView("bootstrap", "views/users/new.gtpl"),
	)
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}