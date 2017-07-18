package shared

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func NewDBConn() *sql.DB {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=1234 dbname=blog")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
