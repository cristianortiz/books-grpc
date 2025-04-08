package restbooksserver

import (
	"encoding/json"
	"net/http"

	"github.com/cristianortiz/books-grpc/internal/pkg/service"
)

const SuccessResponseFieldKey = "status"
const ErrorResponseFieldKey = "error"

type BooksHandler struct {
	BookService service.BookService
}

func NewBooksHandler(bookService service.BookService) *BooksHandler {
	return &BooksHandler{BookService: bookService}
}

func (bh *BooksHandler) GetBookList(w http.ResponseWriter, r *http.Request) {
	books, err := bh.BookService.GetAllBooks()
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, books)
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
