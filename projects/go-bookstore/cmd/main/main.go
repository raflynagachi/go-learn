package main

import (
	"log"
	"net/http"

	"github.com/raflynagachi/go-bookstore/pkg/config"
	"github.com/raflynagachi/go-bookstore/pkg/controller"
	"github.com/raflynagachi/go-bookstore/pkg/routes"
	"github.com/raflynagachi/go-bookstore/repository"
	"github.com/raflynagachi/go-bookstore/service"
)

func main() {
	db := config.OpenDB()
	bookRepository := repository.NewBookRepositoryImpl()
	bookService := service.NewBookServiceImpl(bookRepository, db)
	bookController := controller.NewBookControllerImpl(bookService)
	router := routes.NewRouter(bookController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
