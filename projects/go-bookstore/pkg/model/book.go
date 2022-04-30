package model

import (
	"github.com/jinzhu/gorm"
	"github.com/raflynagachi/go-bookstore/pkg/config"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	db := config.OpenDB()
	db.AutoMigrate(&Book{})
}

// var db *gorm.DB
// func (b *Book) Create() *Book {
// 	db.NewRecord(b)
// 	db.Create(&b)
// 	return b
// }

// func Delete(Id int) Book {
// 	var book Book
// 	db.Where("ID=?", Id).Delete(&book)
// 	return book
// }

// func FindById(Id int) Book {
// 	var book Book
// 	db.Where("ID=?", Id).Find(&book)
// 	return book
// }

// func FindAll() []Book {
// 	var books []Book
// 	db.Find(&books)
// 	return books
// }
