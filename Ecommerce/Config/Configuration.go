package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configuration struct holds the database connection information
type Configuration struct {
	PublicHost        string
	Port              string
	DBSource          string
	DBInitial_Catalog string
	DBUser_ID         string
	DBPassword        string
	DBDriver          string
}

//var Env, err = GetConfiguration()

// GetConfiguration returns a Configuration object with the database connection information
func GetConfiguration() (Configuration, error) {
	if err := godotenv.Load("D:\\Rahul\\Go_Learning\\golang_api\\Ecommerce\\.env"); err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Println("Error: .env file not found")
			return Configuration{}, err
		} else {
			log.Println("Error loading .env file:", err)
			return Configuration{}, err
		}
	}

	return Configuration{
		PublicHost:        getEnv("PUBLIC_HOST", "public_host"),
		Port:              getEnv("PORT", "port"),
		DBSource:          getEnv("DB_Source", "db_source"),
		DBInitial_Catalog: getEnv("DB_Initial_Catalog", "catalog"),
		DBUser_ID:         getEnv("USER", "user"),
		DBPassword:        getEnv("DB_PASSWORD", "password"),
		DBDriver:          getEnv("DRIVER", "sqlserver"),
	}, nil
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
