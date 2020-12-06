package main

import (
	"../../database"
	"../api"
	"log"
)

func main() {
	db, err := database.GetDatabaseHandler(database.MONGODB, "localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = api.RunApi("localhost:8080", db)
	if err != nil {
		log.Fatal(err)
	}
}
