package service

import "github.com/raflynagachi/go-bookstore/pkg/model"

type BookService interface {
	Create(book model.Book) model.Book
	Update(book model.Book, bookId int) model.Book
	Delete(bookId int) model.Book
	FindById(bookId int) model.Book
	FindAll() []model.Book
}
