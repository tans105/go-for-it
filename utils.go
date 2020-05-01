package main

import (
	"net/http"
	"strings"
)

func isUserExists(u User) User {
	var dbUser User
	db.Debug().Select("email,password").Where("email = ?", u.Email).First(&User{}).Scan(&dbUser)
	return dbUser
}

func validatePayload(u User, isLogin bool) string {
	var message string
	if isNullOrEmpty(u.Email) {
		message = EMAIL_NOT_FOUND + SEPARATOR
	}

	if isNullOrEmpty(u.Password) {
		message += PASSWORD_NOT_FOUND + SEPARATOR
	}

	if !isLogin && isNullOrEmpty(u.Name) {
		message += NAME_NOT_FOUND + SEPARATOR
	}

	return message
}

func isNullOrEmpty(s string) bool {
	return len(strings.Trim(s, " ")) == 0 || s == ""
}

func alreadyLoggedIn(req *http.Request) bool {
	var userSession Session
	var count int
	ok := false

	c, err := req.Cookie("uuid")
	if err != nil {
		return false
	}

	db.Debug().Select("email,session_id").Where("session_id = ?", c.Value).First(&Session{}).Scan(&userSession)
	if (Session{}) != userSession {
		db.Debug().Where("email = ?", userSession.Email).Find(&User{}).Count(&count)
		ok = count > 0
	}

	return ok
}

func isValidLogin(u User) bool {
	dbUser := isUserExists(u)

	if (User{}) != dbUser {
		dbPass := dbUser.Password
		return dbPass == u.Password
	} else {
		return false
	}
}
