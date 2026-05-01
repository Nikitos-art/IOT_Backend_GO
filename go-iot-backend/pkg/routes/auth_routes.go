package routes

import (
    "github.com/gorilla/mux"
    "github.com/Nikitos-art/go-iot-backend/pkg/controllers"
    "github.com/Nikitos-art/go-iot-backend/pkg/services"
    "github.com/jinzhu/gorm"
)

func RegisterAuthRoutes(r *mux.Router, db *gorm.DB) {
    authService := services.NewAuthService(db)
    authController := controllers.NewAuthController(authService)

    r.HandleFunc("/register", authController.Register).Methods("POST")
    r.HandleFunc("/login", authController.Login).Methods("POST")
}