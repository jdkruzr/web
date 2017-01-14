package views

import "html/template"

type View struct {
	Template *template.Template
	Layout	string
}

func NewView(layout string, files ...string) *View {
	files = append(files, 
		"views/layouts/footer.gtpl",
		"views/layouts/navbar.gtpl",
		"views/layouts/bootstrap.gtpl",)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	
	return &View{
		Template: 	t,
		Layout:		layout,
	}
}