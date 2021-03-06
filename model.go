package main

import "github.com/jinzhu/gorm"

type DbConfiguration struct {
	Vendor string
	Host string
	Port int
	Username string
	Password string
}

type Response struct {
	Success bool
	Message string
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string
	Mobile   string
}

type Session struct {
	gorm.Model
	Email     string `gorm:"unique;not null"`
	SessionId string `gorm:"unique;not null"`
}
