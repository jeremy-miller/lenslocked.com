package controllers

import "github.com/jeremy-miller/lenslocked.com/views"

type Static struct {
	Home    *views.View
	Contact *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.New("bootstrap", "static/home"),
		Contact: views.New("bootstrap", "static/contact"),
	}
}
