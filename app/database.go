package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang_restful_api")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
