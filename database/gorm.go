package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() (gormDb *gorm.DB) {
	var err error
	if gormDb, err = gorm.Open(postgres.New(postgres.Config{Conn: SqlDb}), &gorm.Config{}); err != nil {
		log.Fatal(err)
		return
	}
	DB = gormDb
	return
}
