package main

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/config"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/routes"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/models"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/middleware"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("no .env file found")
	}

	config.Connect()

	// IMPORTANT: run migrations AFTER DB is ready
	// config.GetDB().AutoMigrate(&models.Book{})
	config.GetDB().AutoMigrate(&models.User{})
	config.GetDB().AutoMigrate(&models.Device{})
	config.GetDB().AutoMigrate(&models.DeviceData{})

	r := mux.NewRouter()

	// 🔥 Logger middleware
	r.Use(middleware.Logger)

	// 🔥 CORS middleware (needed for frontend)
	// r.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// 		w.Header().Set("Access-Control-Allow-Origin", "*")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// 		if r.Method == "OPTIONS" {
	// 			w.WriteHeader(http.StatusOK)
	// 			return
	// 		}

	// 		next.ServeHTTP(w, r)
	// 	})
	// })

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