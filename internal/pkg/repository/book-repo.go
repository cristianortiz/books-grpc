package repository

import (
	"github.com/cristianortiz/books-grpc/internal/pkg/model"
	"gorm.io/gorm"
)

// represents a repo for managing book resources, it contains *gorm.DB typ field
// for interacting with database, implementing CRUD operations trough it
type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) AddBook(book *model.DBBook) error {
	result := br.db.Create(book)
	if result.Error != nil {
		// Devolver el error si la operación falla
		return result.Error
	}
	return nil
}
func (br *BookRepository) UpdateBook(book *model.DBBook) {
	br.db.Model(&book).Where("isbn=?", book.Isbn).Update("name", "publisher")
}

func (br *BookRepository) GetBook(isbn int) *model.DBBook {
	var book model.DBBook
	br.db.First(&book, isbn)
	return &book
}
func (br *BookRepository) GetAllBooks() ([]*model.DBBook, error) {
	books := make([]*model.DBBook, 0)
	//var books []*model.DBBook
	err := br.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (br *BookRepository) RemoveBook(isbn int) {
	br.db.Delete(&model.DBBook{}, isbn)
}
