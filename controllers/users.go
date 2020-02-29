package controllers

import (
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

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
