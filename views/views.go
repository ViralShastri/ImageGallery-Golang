// Package views for creating new Views
package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	// LayoutDir is a Layout Directory
	LayoutDir string = "views/layouts/"
	// TemplateExt is a Template Extension
	TemplateExt string = ".gohtml"
	// TemplateDir is a Template Directory
	TemplateDir = "views/"
)

// NewView Function
func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)

	template, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{*template, layout}
}

// View Structure for creating view
type View struct {
	Template template.Template
	Layout   string
}

func (view *View) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := view.Render(rw, nil); err != nil {
		panic(err)
	}
}

// Render Method to Execute a View
func (view *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return view.Template.ExecuteTemplate(w, view.Layout, data)
}

func layoutFiles() []string {
	layouts, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return layouts
}

// addTemplatePath takes in a slice of strings
// representing files paths for templates, and
// its prepends TemplateDir directory to each
// in the slice
//
// Eg: input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for index, file := range files {
		files[index] = TemplateDir + file
	}
}

// addTemplateExt takes in a slice of strings
// representing files paths for templates, and
// its postpends TemplateExt Extension to each
// in the slice
//
// Eg: input {"home"} would result in the output
// {"views.gohtml"} if TemplateDir == ".gohtml"
func addTemplateExt(files []string) {
	for index, file := range files {
		files[index] = file + TemplateExt
	}
}
