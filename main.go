package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)

	http.ListenAndServe(port, nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func signup(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
