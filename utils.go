package main

import (
	"net/http"
	"strings"
)

func isUserExists(u User) bool {
	_, present := dbUsers[u.Email]
	return present
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
	c, err := req.Cookie("uuid")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

func isValidLogin(u User) bool {
	//TODO: Store encrypted password and compare with the same
	if isUserExists(u) {
		dbPass := dbUsers[u.Email].Password
		return dbPass == u.Password
	} else {
		return false
	}
}
