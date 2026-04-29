package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost port=5432 user=apollo_26 dbname=go_rest sslmode=disable password=Eigentlich&3"

	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}