package main

import grpcbooksserver "github.com/cristianortiz/books-grpc/internal/grpc-books-server"

func main() {
	app := grpcbooksserver.NewApp()
	app.Start()
	app.Shutdown()
}
