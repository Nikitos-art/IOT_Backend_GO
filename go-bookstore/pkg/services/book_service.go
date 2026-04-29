package services

import "github.com/Nikitos-art/go-bookstore/pkg/models"

func CreateBook(b *models.Book) (*models.Book, error) {
	err := b.CreateBook()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllBooks() ([]models.Book, error) {
	return models.GetAllBooks()
}

func GetBookById(id int64) (*models.Book, error) {
	return models.GetBookById(id)
}

func DeleteBook(id int64) error {
	return models.DeleteBook(id)
}

func UpdateBook(b *models.Book) error {
	return models.UpdateBook(b)
}