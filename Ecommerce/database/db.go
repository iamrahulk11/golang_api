package database

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/golang_api/Ecommerce/Config"
)

func NewSqlStorge(cfg *config.Configuration) (*sql.DB, error) {
	connStr := fmt.Sprintf("Data Source=%s;Initial Catalog=%s;User ID=%s;Password=%s;TrustServerCertificate=True;Connect Timeout=30;", cfg.DBSource, cfg.DBInitial_Catalog, cfg.DBUser_ID, cfg.DBPassword)
	db, err := sql.Open("sqlserver", connStr)

	if err != nil {
		log.Fatal()
	}
	return db, nil
}
