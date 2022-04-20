package main

import (
	_ "github.com/go-sql-driver/mysql"

	"gorestful/app"
	"gorestful/controller"
	"gorestful/helper"
	"gorestful/middleware"
	"gorestful/repository"
	"gorestful/service"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
