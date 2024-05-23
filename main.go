package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/thomzes/go-restful/helper"
	"github.com/thomzes/go-restful/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
