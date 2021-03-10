package controllers

import (
	"net/http"

	"github.com/ViralShastri/usegolang/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// Users Type
type Users struct {
	NewView *views.View
}

// New is a Users type method
// To Render a Users View
func (users *Users) New(rw http.ResponseWriter, r *http.Request) {
	if err := users.NewView.Render(rw, nil); err != nil {
		panic(err)
	}
}
