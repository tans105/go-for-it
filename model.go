package main

import "github.com/jinzhu/gorm"

type Response struct {
	Success bool
	Message string
}

type User struct {
	gorm.Model
	Email    string
	Password string
	Name     string
	Mobile   string
}

type Session struct {
	gorm.Model
	SessionId string
}
