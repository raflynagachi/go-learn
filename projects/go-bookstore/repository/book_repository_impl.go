package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/raflynagachi/go-bookstore/pkg/model"
)

type BookRepositoryImpl struct {
}

func NewBookRepositoryImpl() *BookRepositoryImpl {
	return &BookRepositoryImpl{}
}

func (b *BookRepositoryImpl) Create(db *gorm.DB, book model.Book) model.Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func (b *BookRepositoryImpl) Update(db *gorm.DB, book model.Book, bookId int) model.Book {
	var newBook model.Book
	db.Where("ID=?", bookId).Find(&newBook)
	db.Model(&newBook).Updates(book)
	return newBook
}

func (b *BookRepositoryImpl) Delete(db *gorm.DB, Id int) model.Book {
	var book model.Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}

func (b *BookRepositoryImpl) FindById(db *gorm.DB, Id int) model.Book {
	var book model.Book
	db.Where("ID=?", Id).Find(&book)
	return book
}

func (b *BookRepositoryImpl) FindAll(db *gorm.DB) []model.Book {
	var books []model.Book
	db.Find(&books)
	return books
}
