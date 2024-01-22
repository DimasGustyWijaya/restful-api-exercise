package app

import (
	"database/sql"
	"restful-api/helper"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/tokopadui")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxOpenConns(5)
	db.SetConnMaxIdleTime(30 * time.Minute)

	return db
}
