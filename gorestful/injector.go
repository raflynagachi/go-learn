//go:build wireinject
// +build wireinject

package main

import (
	"gorestful/app"
	"gorestful/controller"
	"gorestful/middleware"
	"gorestful/repository"
	"gorestful/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

var newRouter = wire.NewSet(
	app.NewRouter,
	wire.Bind(new(http.Handler), new(*httprouter.Router)),
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		newRouter,
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
