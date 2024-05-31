package main

import (
	"library/config"
	"library/database"
	"log"
)

func main() {
	gormDb := database.Open()
	defer func() {
		if err := database.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	config.Router(gormDb)
}
