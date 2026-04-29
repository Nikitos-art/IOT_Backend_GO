package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Nikitos-art/go-bookstore/pkg/config"
)

type Book struct {
	gorm.Model

	Name        string `json:"name" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Publication string `json:"publication" validate:"required"`
}

func (b *Book) CreateBook() *Book {
	config.GetDB().Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	config.GetDB().Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := config.GetDB().Where("id = ?", id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	config.GetDB().Where("id = ?", id).Delete(&book)
	return book
}