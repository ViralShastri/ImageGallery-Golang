// Package views for creating new Views
package views

import "html/template"

// NewView Function
func NewView(files ...string) *View {
	files = append(files, "./layouts/footer.gohtml")
	template, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{*template}
}

// View Structure for creating view
type View struct {
	Template template.Template
}
