package main

import (
	"database/sql"
	"log"

	config "github.com/golang_api/Ecommerce/Config"
	"github.com/golang_api/Ecommerce/cmd/api"
	"github.com/golang_api/Ecommerce/database"
)

func main() {
	Configuration := config.GetConfiguration()
	db, err := database.NewSqlStorge(&Configuration)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: connected")
}
