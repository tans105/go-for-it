package main

func isUserExists(u User) bool {
	_, present := dbUser[u.email]
	return present
}

func validatePayload(u User) string {
	var message string
	if isNullOrEmpty(u.email) {
		message = EMAIL_NOT_FOUND + SEPARATOR
	}

	if isNullOrEmpty(u.mobile) {
		message += MOBILE_NOT_FOUND + SEPARATOR
	}

	if isNullOrEmpty(u.name) {
		message += NAME_NOT_FOUND + SEPARATOR
	}

	if isNullOrEmpty(u.password) {
		message += PASSWORD_NOT_FOUND + SEPARATOR
	}

	return message
}

func isNullOrEmpty(s string) bool {
	return len(s) == 0 || s == ""
}
