package restbooksserver

import "github.com/cristianortiz/books-grpc/internal/pkg/model"

// these functions serve  as mapping utilities for converting between
// Book and DBook data models this convertion will be usefull across different parts of the app
// centralice the conversion logic
func DBBook(book *model.Book) *model.DBBook {
	dbBook := &model.DBBook{
		Isbn:      book.Isbn,
		Name:      book.Name,
		Publisher: book.Publisher,
	}

	return dbBook
}
func Book(dbBooks *model.DBBook) *model.Book {
	book := &model.Book{
		Isbn:      dbBooks.Isbn,
		Name:      dbBooks.Name,
		Publisher: dbBooks.Publisher,
	}

	return book
}
