// package models

// import (
// 	"github.com/jinzhu/gorm"
// 	"github.com/Nikitos-art/go-bookstore/pkg/config"
// )

// type Book struct {
// 	gorm.Model

// 	Name        string `json:"name" validate:"required"`
// 	Author      string `json:"author" validate:"required"`
// 	Publication string `json:"publication" validate:"required"`

// 	UserID uint `json:"user_id"`
// }
// func (b *Book) CreateBook() error {
// 	result := config.GetDB().Create(b)
// 	return result.Error
// }

// func GetAllBooks() ([]Book, error) {
// 	var books []Book
// 	result := config.GetDB().Find(&books)
// 	return books, result.Error
// }

// func GetBookById(id int64) (*Book, error) {
// 	var book Book
// 	result := config.GetDB().Where("id = ?", id).Find(&book)
// 	return &book, result.Error
// }

// func DeleteBook(id int64) error {
// 	result := config.GetDB().Where("id = ?", id).Delete(&Book{})
// 	return result.Error
// }

// func UpdateBook(b *Book) error {
// 	result := config.GetDB().Save(b)
// 	return result.Error
// }