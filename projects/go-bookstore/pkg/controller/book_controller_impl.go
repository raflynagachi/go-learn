package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raflynagachi/go-bookstore/pkg/model"
	"github.com/raflynagachi/go-bookstore/pkg/utils"
	"github.com/raflynagachi/go-bookstore/service"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookControllerImpl(bookService service.BookService) *BookControllerImpl {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (b *BookControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	newBook := &model.Book{}
	utils.ParseBody(r, newBook)

	book := b.BookService.Create(*newBook)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func (b *BookControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	newBook := &model.Book{}
	utils.ParseBody(r, newBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}

	book := b.BookService.Update(*newBook, ID)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func (b *BookControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}

	book := b.BookService.Delete(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func (b *BookControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}

	book := b.BookService.FindById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func (b *BookControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	books := b.BookService.FindAll()
	res, _ := json.Marshal(books)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}
