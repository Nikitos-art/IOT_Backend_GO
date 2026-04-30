package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Nikitos-art/go-bookstore/pkg/middleware"
	"github.com/Nikitos-art/go-bookstore/pkg/models"
	"github.com/Nikitos-art/go-bookstore/pkg/services"
	"github.com/Nikitos-art/go-bookstore/pkg/utils"
)

type BookController struct {
	service *services.BookService
}

func NewBookController(s *services.BookService) *BookController {
	return &BookController{
		service: s,
	}
}

func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserKey)

	if userID == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	uid := userID.(uint)

	newBooks, err := c.service.GetBooksByUser(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(newBooks)
}

func (c *BookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	bookDetails, err := c.service.GetBookById(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(bookDetails)
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}

	// parse request
	if err := utils.ParseBody(r, newBook); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := utils.ValidateStruct(newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 🔥 GET USER ID FROM JWT
	userID := r.Context().Value(middleware.UserKey)
	if userID == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	newBook.UserID = userID.(uint)

	// save book
	created, err := c.service.CreateBook(newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(created)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	err = c.service.DeleteBook(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "book deleted",
	})
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}

	if err := utils.ParseBody(r, updateBook); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	bookDetails, err := c.service.GetBookById(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	err = c.service.UpdateBook(bookDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(bookDetails)
}