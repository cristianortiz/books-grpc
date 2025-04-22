package grpcbooksserver

import (
	"context"
	"fmt"
	"log"

	"github.com/cristianortiz/books-grpc/internal/pkg/model"
	"github.com/cristianortiz/books-grpc/internal/pkg/proto"
)

func (a *App) AddBook(_ context.Context, req *proto.Book) (*proto.AddBookResponse, error) {
	log.Println("adding book")

	book := &model.DBBook{
		Isbn:      int(req.Isbn),
		Name:      req.Name,
		Publisher: req.Publisher,
	}

	// Intentar agregar el libro a la base de datos
	err := a.bookRepo.AddBook(book)
	if err != nil {
		// Loguear el error y devolverlo al cliente gRPC
		log.Printf("failed to add book: %v", err)
		return nil, fmt.Errorf("failed to add book: %w", err)
	}
	return &proto.AddBookResponse{Status: fmt.Sprintf("book with isbn(%d), name(%s), publisher (%s) added succesfully", book.Isbn, book.Name, book.Publisher)}, nil
}
