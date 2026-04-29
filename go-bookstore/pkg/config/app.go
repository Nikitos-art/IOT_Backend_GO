package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// fmt.Println("DEBUG DB_HOST:", host)
	// fmt.Println("DEBUG DB_PORT:", port)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, name, port,
	)

	// fmt.Println("DEBUG DSN:", dsn)

	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}