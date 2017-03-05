package controllers

import "github.com/jdkruzr/web/views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("bootstrap", "static/home"), // We used to have "views" and "gtpl" here, recall
		Contact: views.NewView("bootstrap", "static/contact"),
		Faq: views.NewView("bootstrap", "static/faq"),
	}
}

type Static struct {
	Home	*views.View
	Contact	*views.View
  Faq *views.View
}

