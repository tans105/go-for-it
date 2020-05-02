package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/url"
	"os"
)

var db *gorm.DB

func initDb() {
	conn, err := gorm.Open("postgres", getDbConfig())

	if err != nil {
		panic(err)
	}
	db = conn
	initDbImport()
}

func getDbConfig() string {
	var conf DbConfiguration = getConfigurationFromJson()
	dsn := url.URL{
		User:     url.UserPassword(conf.Username, conf.Password),
		Scheme:   conf.Vendor,
		Host:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	return dsn.String()
}

func initDbImport() {
	db.Debug().AutoMigrate(&User{}, &Session{})
}

func getConfigurationFromJson() DbConfiguration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := DbConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Unable to fetch configuration:", err)
	}
	return configuration
}
