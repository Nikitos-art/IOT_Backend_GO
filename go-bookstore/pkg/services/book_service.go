package services

import (
	"github.com/Nikitos-art/go-bookstore/pkg/models"
	"github.com/jinzhu/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{db: db}
}

func (s *BookService) CreateBook(b *models.Book) (*models.Book, error) {
	err := s.db.Create(b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *BookService) GetBooksByUser(userID uint) ([]models.Book, error) {
	var books []models.Book

	err := s.db.Where("user_id = ?", userID).Find(&books).Error
	return books, err
}

func (s *BookService) GetBookById(id int64) (*models.Book, error) {
	return models.GetBookById(id)
}

func (s *BookService) DeleteBook(id int64) error {
	return models.DeleteBook(id)
}

func (s *BookService) UpdateBook(b *models.Book) error {
	return models.UpdateBook(b)
}