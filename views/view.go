package views

import "html/template"
import "path/filepath"
import "net/http"

// Setting some standards for views processing 
var ViewsDir string = "views"
var ViewsExt string = ".gtpl"
var LayoutDir string = ViewsDir + "/layouts"

// View = gtpl file plus whatever layout we're using (currently
// Bootstrap with a navbar and a footer)
type View struct {
	Template *template.Template
	Layout	string
}

func addViewsDirPrefix(files []string) {
	for i, f := range files {
		files[i] = ViewsDir + "/" + f
	}
}

func addViewsExtSuffix(files []string) {
	for i, f := range files {
		files[i] = f + ViewsExt
	}
}

//func (v *View) Render(w http.ResponseWriter, data interface{}) error {
//	return v.Template.ExecuteTemplate(w, v.Layout, data)
//}

// Build a glob of all the template files for presentation to 
// the View "constructor"
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*" + ViewsExt)
	if err != nil {
		panic(err)
	}
	return files
}

// View "constructor" attaches all our layout files to our view
func NewView(layout string, files ...string) *View {
	addViewsDirPrefix(files)
	addViewsExtSuffix(files)
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

// DRY up the Static handler
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

