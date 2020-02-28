package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDir   = "views/layouts/"
	templateExt = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func New(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{Template: t, Layout: layout}
}

// layoutFiles returns a slice of strings representing the layout files used in our application.
func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}
