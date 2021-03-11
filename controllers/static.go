package controllers

import "github.com/ViralShastri/usegolang/views"

// NewStatic is used to create a new Static controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

// Static Type for Static Views
type Static struct {
	Home    *views.View
	Contact *views.View
}
