package app

import (
	"database/sql"
	"time"

	"github.com/thomzes/go-restful/helper"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:kaliwuluh11@tcp(localhost:3306)/golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns((20))
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}
