//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"github.com/thomzes/go-restful/app"
	"github.com/thomzes/go-restful/controller"
	"github.com/thomzes/go-restful/middleware"
	"github.com/thomzes/go-restful/repository"
	"github.com/thomzes/go-restful/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

// ProvideValidatorOptions provides an empty slice of validator options.
func ProvideValidatorOptions() []validator.Option {
	return []validator.Option{}
}

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDb,
		ProvideValidatorOptions,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
