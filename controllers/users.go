package controllers

import (
	"fmt"
	"net/http"

	"github.com/jeremy-miller/lenslocked.com/views"
)

type Users struct {
	NewView *views.View
}

// NewUsers is used to create a newUsers controller. This function will panic if the
// templates are not parsed correctly and should only be used during initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.New("bootstrap", "views/users/new.gohtml"),
	}
}

// New is used to render the form where a user can create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user submits it. This is used to create
// a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a temporary response.")
}
