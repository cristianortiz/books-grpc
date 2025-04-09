package restbooksserver

import (
	"github.com/cristianortiz/books-grpc/internal/pkg/service"
	"github.com/gorilla/mux"
)

func ProvideRouter(bookService service.BookService) *mux.Router {
	r := mux.NewRouter()
	booksHandler := NewBooksHandler(bookService)
	r.HandleFunc("/books", booksHandler.GetBookList).Methods("GET")
	r.HandleFunc("/books", booksHandler.UpsertBookHandler).Methods("POST", "PUT")
	r.HandleFunc("/books/{isbn:[0-9]+}", booksHandler.GetOrRemoveBookHandler).Methods("GET", "DELETE")

	return r
}
