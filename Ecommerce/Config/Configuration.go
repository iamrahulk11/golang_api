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

var Env = GetConfiguration()

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// GetConfiguration returns a Configuration object with the database connection information
func GetConfiguration() Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Configuration{
		PublicHost:        getEnv("PUBLIC_HOST", "http://localhost"),
		Port:              getEnv("PORT", "8080"),
		DBSource:          getEnv("DB_Source", "source"),
		DBInitial_Catalog: getEnv("DB_Initial_Catalog", "catalog"),
		DBUser_ID:         getEnv("USER", "dbuser"),
		DBDriver:          getEnv("DRIVER", "sqlserver"),
		DBPassword:        getEnv("DB_PASSWORD", "myPassword"),
	}
}
