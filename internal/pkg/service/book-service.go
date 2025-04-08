package service

import (
	"fmt"

	"github.com/cristianortiz/books-grpc/internal/pkg/model"
	"github.com/cristianortiz/books-grpc/internal/pkg/repository"
)

type BookService struct {
	booksRepo *repository.BookRepository
}

func NewBooksService(booksRepo *repository.BookRepository) BookService {
	return BookService{booksRepo: booksRepo}
}
func (bs *BookService) AddBook(book *model.DBBook) {
	bs.booksRepo.AddBook(book)
}

func (bs *BookService) GetBook(isbn int) (*model.DBBook, error) {
	book := bs.booksRepo.GetBook(isbn)
	if book != nil {
		return book, nil
	}
	return nil, fmt.Errorf("book with isbn %d was not found", isbn)
}

func (bs *BookService) GetAllBooks() ([]*model.DBBook, error) {
	books, err := bs.booksRepo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, fmt.Errorf("no books founded")
	}
	return books, nil
}

func (bs *BookService) RemoveBook(isbn int) {
	bs.booksRepo.RemoveBook(isbn)
}

func (bs *BookService) UpdateBook(book *model.DBBook) {
	bs.booksRepo.UpdateBook(book)
}
