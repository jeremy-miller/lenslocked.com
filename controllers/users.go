package controllers

import (
	"fmt"
	"net/http"

	"github.com/jeremy-miller/lenslocked.com/models"
	"github.com/jeremy-miller/lenslocked.com/views"
)

type Users struct {
	NewView     *views.View
	userService *models.UserService
}

// NewUsers is used to create a newUsers controller. This function will panic if the
// templates are not parsed correctly and should only be used during initial setup.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:     views.New("bootstrap", "users/new"),
		userService: us,
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

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create is used to process the signup form when a user submits it. This is used to create
// a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{Name: form.Name, Email: form.Email}
	if err := u.userService.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, form)
}
