package service

import (
	"github.com/jinzhu/gorm"
	"github.com/raflynagachi/go-bookstore/pkg/model"
	"github.com/raflynagachi/go-bookstore/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *gorm.DB
}

func NewBookServiceImpl(bookRepository repository.BookRepository, db *gorm.DB) *BookServiceImpl {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             db,
	}
}

func (bookService *BookServiceImpl) Create(book model.Book) model.Book {
	newBook := bookService.BookRepository.Create(bookService.DB, book)
	return newBook
}

func (bookService *BookServiceImpl) Update(book model.Book, bookId int) model.Book {
	newBook := bookService.BookRepository.Update(bookService.DB, book, bookId)
	return newBook
}

func (bookService *BookServiceImpl) Delete(bookId int) model.Book {
	book := bookService.BookRepository.Delete(bookService.DB, bookId)
	return book
}

func (bookService *BookServiceImpl) FindById(bookId int) model.Book {
	book := bookService.BookRepository.FindById(bookService.DB, bookId)
	return book
}

func (bookService *BookServiceImpl) FindAll() []model.Book {
	books := bookService.BookRepository.FindAll(bookService.DB)
	return books
}
