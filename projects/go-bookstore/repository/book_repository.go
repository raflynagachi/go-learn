package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/raflynagachi/go-bookstore/pkg/model"
)

type BookRepository interface {
	Create(db *gorm.DB, book model.Book) model.Book
	Update(db *gorm.DB, book model.Book, bookId int) model.Book
	Delete(db *gorm.DB, Id int) model.Book
	FindById(db *gorm.DB, Id int) model.Book
	FindAll(db *gorm.DB) []model.Book
}
