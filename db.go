package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/url"
)

var db *gorm.DB

func initDb() {
	conn, err := gorm.Open("postgres", getDbConfig())

	if err != nil {
		panic(err)
	}
	db = conn
	setup()
}

func getDbConfig() string {
	dsn := url.URL{
		User:     url.UserPassword("postgres", "postgres"),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", "localhost", 5432),
		Path:     "postgres",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	return dsn.String()
}

func setup() {
	db.Debug().AutoMigrate(&User{} , &Session{})
}
