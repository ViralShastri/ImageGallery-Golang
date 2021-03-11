package controllers

import (
	"fmt"
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

// SignupForm Type
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// New is used to render a form
// where New User can create a new account
//
// GET /signup
func (users *Users) New(rw http.ResponseWriter, r *http.Request) {
	if err := users.NewView.Render(rw, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form
// when a user submit it.
// This is used to create a new user account
//
// POST /signup
func (users *Users) Create(rw http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(rw, form)
}
