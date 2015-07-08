package controllers

import (
	"html/template"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, struct{ Name string }{Name: "World"})
}
