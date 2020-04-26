package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var dbSessions = map[string]string{}
var dbUsers = map[string]User{}

type User struct {
	email    string
	password string
	name     string
	mobile   string
}

func init() {
	tpl = template.Must(template.ParseGlob("public/templates/*html"))
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))

	http.HandleFunc("/", login)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/home", home)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(PORT, nil)
}
