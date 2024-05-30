package main

import (
	"library/database"
	"log"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// if err := router.Run(db); err != nil {
	// 	log.Fatal(err)
	// }
}