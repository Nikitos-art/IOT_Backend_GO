package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Nikitos-art/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	// API sanity check route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bookstore API running"))
	}).Methods("GET")

	// Subrouter for /book
	bookRouter := router.PathPrefix("/book").Subrouter()

	bookRouter.HandleFunc("", controllers.GetBook).Methods("GET")
	bookRouter.HandleFunc("", controllers.CreateBook).Methods("POST")
	bookRouter.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")
	bookRouter.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")
	bookRouter.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")
}