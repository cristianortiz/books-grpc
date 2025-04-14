package main

import restbooksserver "github.com/cristianortiz/books-grpc/internal/rest-books-server"

func main() {
	app := restbooksserver.NewApp()
	app.Start()
	app.Shutdown()
}
