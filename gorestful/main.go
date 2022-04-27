package main

import (
	_ "github.com/go-sql-driver/mysql"

	"gorestful/helper"
	"gorestful/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {
	// validate := validator.New()
	// db := app.NewDB()

	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)

	// router := app.NewRouter(categoryController)
	// authMiddleware := middleware.NewAuthMiddleware(router)

	server := InitializeServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
