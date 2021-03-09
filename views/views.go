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
)

// NewView Function
func NewView(layout string, files ...string) *View {

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

// Render Method to Execute a View
func (view *View) Render(w http.ResponseWriter, data interface{}) error {
	return view.Template.ExecuteTemplate(w, view.Layout, data)
}

func layoutFiles() []string {
	layouts, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return layouts
}
