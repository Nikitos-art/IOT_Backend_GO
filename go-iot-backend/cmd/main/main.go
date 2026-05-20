// The entrypoint where:
// env vars are loaded
// 3 db tables are autoloaded
// routes are registered
// and middleware all come together in an orchestra like manner
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/config"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/middleware"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/models"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("no .env file found")
	}

	config.Connect()

	// IMPORTANT: run migrations AFTER DB is ready
	config.GetDB().AutoMigrate(&models.User{})
	config.GetDB().AutoMigrate(&models.Device{})
	config.GetDB().AutoMigrate(&models.DeviceData{})

	r := mux.NewRouter()

	r.Use(middleware.Logger)

	routes.RegisterIoTBackendRoutes(r)
	routes.RegisterAuthRoutes(r, config.GetDB())

	fmt.Println(`
	========================================
	🚀  IoT BACKEND SERVER RUNNING
	========================================
	`)
	fmt.Printf("👉 URL: http://localhost%s\n", ":9010")
	fmt.Println("👉 Click or open in browser")
	fmt.Println("========================================")

	log.Fatal(http.ListenAndServe(":9010", r))

}
