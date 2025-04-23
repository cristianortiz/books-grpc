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
	book, error := bs.booksRepo.GetBook(isbn)
	if error != nil {
		return nil, fmt.Errorf("book with isbn %d was not found", isbn)
	}
	return book, nil
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
