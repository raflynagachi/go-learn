package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/raflynagachi/go-bookstore/pkg/utils"
)

func OpenDB() *gorm.DB {
	db, err := gorm.Open("mysql", "adminer:password@tcp(localhost:3306)/go-bookstore?charset=utf8&parseTime=True")
	utils.PanicIfError(err)
	return db
}
