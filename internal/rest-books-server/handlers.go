package restbooksserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cristianortiz/books-grpc/internal/pkg/model"
	"github.com/cristianortiz/books-grpc/internal/pkg/service"
	"github.com/gorilla/mux"
)

const SuccessResponseFieldKey = "status"
const ErrorResponseFieldKey = "error"

type BooksHandler struct {
	bookService service.BookService
}

// constructor fucntion that creates a new BooksHandler instance with the provided bookService
func NewBooksHandler(bookService service.BookService) *BooksHandler {
	return &BooksHandler{bookService: bookService}
}

func (bh *BooksHandler) GetBookList(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookService.GetAllBooks()
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, books)
}

func (bh *BooksHandler) GetOrRemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	muxVar := mux.Vars(r)
	isbnStr := muxVar["isbn"]
	isbn, err := strconv.Atoi(isbnStr)
	if err != nil || isbn == 0 {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if r.Method == "GET" {
		book, err := bh.bookService.GetBook(isbn)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, Book(book))
	}
	if r.Method == "DELETE" {
		bh.bookService.RemoveBook(isbn)
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "book removed"})
	}
}

func (bh *BooksHandler) UpsertBookHandler(w http.ResponseWriter, r *http.Request) {
	var book *model.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if r.Method != "POST" && r.Method != "PUT" {
		respondWithError(w, http.StatusMethodNotAllowed, "invalid request method")
		return
	}

	if r.Method == "POST" {
		bh.bookService.AddBook(DBBook(book))
	}

	if r.Method == "PUT" {
		bh.bookService.UpdateBook(DBBook(book))
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "book removed"})
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "book upserted"})

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{SuccessResponseFieldKey: message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
