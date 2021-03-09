// Package views for creating new Views
package views

import "html/template"

// NewView Function
func NewView(layout string, files ...string) *View {
	files = append(files, "views/layouts/navbar.gohtml", "views/layouts/bootstrap.gohtml")
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
