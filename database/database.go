package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var SqlDb *sql.DB

func init() {
	var err error
	if SqlDb, err = sql.Open("postgres","postgresql://alfatinfernanda:@localhost:5432/library?sslmode=disable"); err != nil {
		log.Fatal(err)
	}
}