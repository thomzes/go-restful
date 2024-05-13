package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/thomzes/go-restful/app"
	"github.com/thomzes/go-restful/controller"
	"github.com/thomzes/go-restful/helper"
	"github.com/thomzes/go-restful/middleware"
	"github.com/thomzes/go-restful/repository"
	"github.com/thomzes/go-restful/service"
)

func main() {

	db := app.NewDb()
	validate := validator.New()
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
