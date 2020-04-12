package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
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
	fs := http.FileServer(http.Dir("public/stylesheets"))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", fs))
	http.HandleFunc("/", login)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/home", home)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(PORT, nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	var message string
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	} else {
		if req.Method == http.MethodPost {
			pass := true
			u := User{
				email:    req.FormValue("email"),
				password: req.FormValue("password"),
			}

			message = validatePayload(u, true)
			if len(message) > 0 || !isValidLogin(u) {
				pass = false
			}

			if pass {
				//create session
				uuid := strconv.Itoa(int(rand.Float64()*2089)) + "-" + strconv.Itoa(int(rand.Float64()*9973)) //TODO: Add package to generate session ID
				http.SetCookie(w, &http.Cookie{
					Name:  "uuid",
					Value: uuid,
				})
				dbSessions[uuid] = u.email
				http.Redirect(w, req, "/home", http.StatusSeeOther)
			}
		}
	}

	tpl.ExecuteTemplate(w, "index.html", message)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	}

	var message string
	if req.Method == http.MethodPost {
		u := User{
			email:    req.FormValue("email"),
			password: req.FormValue("password"),
			name:     req.FormValue("name"),
			mobile:   req.FormValue("mobile"),
		}
		if !isUserExists(u) {
			message = validatePayload(u, false)
			if len(message) == 0 {
				dbUsers[u.email] = u
				message = SUCCESSFULL_REGISTRATION
			} else {
				message = COULD_NOT_REGISTER + ":" + message
			}
		} else {
			message = EMAIL_ALREADY_TAKEN
		}
	}

	tpl.ExecuteTemplate(w, "signup.html", message)
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("uuid")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "uuid",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
