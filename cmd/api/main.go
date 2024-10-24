package main

import (
	"github.com/abroudoux/go-rest-api-boilerplate/internal/database"
	"github.com/abroudoux/go-rest-api-boilerplate/internal/router"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.Migrate(db)
	router.StartNewRouter()
}
