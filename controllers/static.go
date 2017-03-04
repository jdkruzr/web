package controllers

import "github.com/jdkruzr/web/views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "views/static/home.gtpl"),
		Contact: views.NewView("bootstrap", "views/static/contact.gtpl"),
		Faq: views.NewView("bootstrap", "views/static/faq.gtpl"),
	}
}

type Static struct {
	Home	*views.View
	Contact	*views.View
  Faq *views.View
}

