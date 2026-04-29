package main

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Nikitos-art/go-bookstore/pkg/config"
	"github.com/Nikitos-art/go-bookstore/pkg/routes"
	"github.com/Nikitos-art/go-bookstore/pkg/models"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("no .env file found")
	}

	config.Connect()

	// IMPORTANT: run migrations AFTER DB is ready
	config.GetDB().AutoMigrate(&models.Book{})

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	fmt.Println(`
	========================================
	🚀  BOOKSTORE SERVER RUNNING
	========================================
	`)
	fmt.Printf("👉 URL: http://localhost%s\n", ":9010")
	fmt.Println("👉 Click or open in browser")
	fmt.Println("========================================")

	log.Fatal(http.ListenAndServe(":9010", r))

}