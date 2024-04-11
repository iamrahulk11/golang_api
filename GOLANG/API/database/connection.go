package database

import (
	"database/sql"
	_config "example/API/config"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" // Import your SQL driver here
)

func CreateConnection(cfg _config.Configuration) (*sql.DB, error) {
	//db, err := sql.Open(driver, connectionString)

	connStr := fmt.Sprintf("Data Source=%s;Initial Catalog=%s;User ID=%s;Password=%s;TrustServerCertificate=True;Connect Timeout=30;", cfg.DBSource, cfg.DBInitial_Catalog, cfg.DBUser_ID, cfg.DBPassword)
	driver := cfg.DBDriver

	// Open a database connection using the DSN
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}

	// Ping the database to check if connection is successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//fmt.Println("Successfully connected to the database!")
	return db, nil
}
