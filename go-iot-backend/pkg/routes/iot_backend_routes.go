package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/controllers"
	"github.com/Nikitos-art/IOT_Backend_GO/go-iot-backend/pkg/middleware"
)

func RegisterIoTBackendRoutes(router *mux.Router) {

	// -------------------------
	// API sanity check route
	// -------------------------
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("IoT Backend is API running"))
	}).Methods("GET")

	// -------------------------
	// BOOKS
	// -------------------------
	// bookService := services.NewBookService(config.GetDB())
	// bookController := controllers.NewBookController(bookService)

	// bookRouter := router.PathPrefix("/books").Subrouter()
	// bookRouter.Use(middleware.AuthMiddleware)

	// bookRouter.HandleFunc("", bookController.GetBooks).Methods("GET")
	// bookRouter.HandleFunc("", bookController.CreateBook).Methods("POST")
	// bookRouter.HandleFunc("/{bookId}", bookController.GetBookById).Methods("GET")
	// bookRouter.HandleFunc("/{bookId}", bookController.UpdateBook).Methods("PUT")
	// bookRouter.HandleFunc("/{bookId}", bookController.DeleteBook).Methods("DELETE")

	// -------------------------
	// DEVICES
	// -------------------------
	deviceRouter := router.PathPrefix("/devices").Subrouter()
	deviceRouter.Use(middleware.AuthMiddleware)

	deviceRouter.HandleFunc("", controllers.CreateDevice).Methods("POST")
	deviceRouter.HandleFunc("", controllers.GetDevicesByUser).Methods("GET")
	deviceRouter.HandleFunc("/{id}", controllers.GetDeviceByID).Methods("GET")
	deviceRouter.HandleFunc("/{id}", controllers.DeleteDevice).Methods("DELETE")

	router.HandleFunc("/ingest", controllers.IngestData).Methods("POST")
}