package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	GormDb *gorm.DB
	SqlDb  *sql.DB
}

func Open() (db DB, err error) {
	db.SqlDb, err = sql.Open("postgres", "postgresql://alfatinfernanda:@localhost:5432/library?sslmode=disable")
	if err != nil {
		return
	}
	if db.GormDb, err = gorm.Open(postgres.New(postgres.Config{Conn: db.SqlDb}), &gorm.Config{}); err != nil {
		return
	}
	return
}