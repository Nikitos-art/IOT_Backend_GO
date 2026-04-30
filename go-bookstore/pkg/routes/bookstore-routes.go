package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Nikitos-art/go-bookstore/pkg/controllers"
	"github.com/Nikitos-art/go-bookstore/pkg/middleware"
	"github.com/Nikitos-art/go-bookstore/pkg/services"
	"github.com/Nikitos-art/go-bookstore/pkg/config"
)

func RegisterBookStoreRoutes(router *mux.Router) {

	// API sanity check route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bookstore API running"))
	}).Methods("GET")

	// create service + controller
	bookService := services.NewBookService(config.GetDB())
	bookController := controllers.NewBookController(bookService)

	// Subrouter for /books
	bookRouter := router.PathPrefix("/books").Subrouter()

	// JWT protection
	bookRouter.Use(middleware.AuthMiddleware)

	// routes
	bookRouter.HandleFunc("", bookController.GetBooks).Methods("GET")
	bookRouter.HandleFunc("", bookController.CreateBook).Methods("POST")
	bookRouter.HandleFunc("/{bookId}", bookController.GetBookById).Methods("GET")
	bookRouter.HandleFunc("/{bookId}", bookController.UpdateBook).Methods("PUT")
	bookRouter.HandleFunc("/{bookId}", bookController.DeleteBook).Methods("DELETE")
}