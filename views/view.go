package views

import "html/template"
import "path/filepath"
import "net/http"

// View = gtpl file plus whatever layout we're using (currently
// Bootstrap with a navbar and a footer)
type View struct {
	Template *template.Template
	Layout	string
}

// What directory does our layout live in?
var LayoutDir string = "views/layouts"

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// Build a glob of all the template files for presentation to 
// the View "constructor"
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gtpl")
	if err != nil {
		panic(err)
	}
	return files
}

// View "constructor" attaches all our layout files to our view
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	
	return &View{
		Template: 	t,
		Layout:		layout,
	}
}