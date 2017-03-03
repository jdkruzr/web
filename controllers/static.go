package controllers

func NewStatic() *Static {
	return &Static{
		HomeView: views.NewView("bootstrap", "views/static/home.gtpl"),
		ContactView: views.NewView("bootstrap", "views/static/contact.gtpl"),
		FAQView: views.NewView("bootstrap", "views/static/faq.gtpl"),
	}
}

type Static struct {
	HomeView	*views.View
	ContactView	*views.View
}

