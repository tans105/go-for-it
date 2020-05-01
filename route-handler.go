package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

func login(w http.ResponseWriter, req *http.Request) {
	var message string
	pass := true
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	} else {
		if req.Method == http.MethodPost {
			u := User{
				Email:    req.FormValue("email"),
				Password: req.FormValue("password"),
			}

			message = validatePayload(u, true)

			if len(message) > 0 {
				pass = false
			} else if !isValidLogin(u) {
				pass = false
				message = INVALID_CREDENTIALS
			}

			if pass {
				//create session
				uuid := strconv.Itoa(int(rand.Float64()*2089)) + "-" + strconv.Itoa(int(rand.Float64()*9973)) //TODO: Add package to generate session ID
				http.SetCookie(w, &http.Cookie{
					Name:  "uuid",
					Value: uuid,
				})

				db.Debug().Delete(Session{}, "email = ?", u.Email)
				db.Create(&Session{
					Email:     u.Email,
					SessionId: uuid,
				})
				http.Redirect(w, req, "/home", http.StatusSeeOther)
			}
		}
	}

	tpl.ExecuteTemplate(w, "index.html", Response{pass, message})
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	}

	var pass bool
	var message string
	if req.Method == http.MethodPost {
		u := User{
			Email:    req.FormValue("email"),
			Password: req.FormValue("password"),
			Name:     req.FormValue("name"),
			Mobile:   req.FormValue("mobile"),
		}

		message = validatePayload(u, false)

		if len(message) == 0 { // request has all the necessary data
			rowsAffected := db.Create(&u).RowsAffected

			if rowsAffected > 0 { // if insert is successful
				pass = true
				message = SUCCESSFULL_REGISTRATION
			} else { // email already present
				pass = false
				message = EMAIL_ALREADY_TAKEN
			}
		} else {
			pass = false
			message = COULD_NOT_REGISTER + ":" + message
		}
	}

	tpl.ExecuteTemplate(w, "signup.html", Response{pass, message})
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

	c = &http.Cookie{ // remove the cookie
		Name:   "uuid",
		Value:  "",
		MaxAge: -1,
	}

	db.Debug().Unscoped().Where("session_id = ?", c.Value).Delete(&Session{})
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
