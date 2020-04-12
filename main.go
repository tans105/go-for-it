package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var dbSession = map[string]string{}
var dbUser = map[string]User{}

type User struct {
	email    string
	password string
	name     string
	mobile   string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)

	http.ListenAndServe(PORT, nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	//TODO: check already logged in
	//TODO: Add package to generate session ID
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func signup(w http.ResponseWriter, req *http.Request) {
	var message string
	if req.Method == http.MethodPost {
		u := User{
			email:    req.FormValue("email"),
			password: req.FormValue("password"),
			name:     req.FormValue("name"),
			mobile:   req.FormValue("mobile"),
		}
		if !isUserExists(u) {
			message = validatePayload(u)
			if len(message) == 0 {
				dbUser[u.email] = u
				message = SUCCESSFULL_REGISTRATION
			} else {
				message = COULD_NOT_REGISTER + ":" + message
			}
		} else {
			message = EMAIL_ALREADY_TAKEN
		}
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", message)
}
