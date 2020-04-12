package main

import (
	"fmt"
	"net/http"
	"strings"
)

func isUserExists(u User) bool {
	_, present := dbUsers[u.email]
	return present
}

func validatePayload(u User, isLogin bool) string {
	var message string
	if isNullOrEmpty(u.email) {
		message = EMAIL_NOT_FOUND + SEPARATOR
	}

	if isNullOrEmpty(u.password) {
		message += PASSWORD_NOT_FOUND + SEPARATOR
	}

	if !isLogin && isNullOrEmpty(u.name) {
		message += NAME_NOT_FOUND + SEPARATOR
	}
	fmt.Println("Length", len(message))
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
		dbPass := dbUsers[u.email].password
		return dbPass == u.password
	} else {
		return false
	}
}
