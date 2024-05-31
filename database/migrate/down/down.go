package main

import (
	"library/database"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)
func main() {
	defer func() {
		if err := database.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	driver, err := postgres.WithInstance(database.SqlDb, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migrateDb, err := migrate.NewWithDatabaseInstance("file://database/migrate/files", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err = migrateDb.Steps(-1); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}