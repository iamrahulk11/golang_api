package config

// Configuration struct holds the database connection information
type Configuration struct {
	DBSource          string
	DBInitial_Catalog string
	DBUser_ID         string
	DBPassword        string
	DBDriver          string
}

// GetConfiguration returns a Configuration object with the database connection information
func GetConfiguration() Configuration {
	return Configuration{
		DBSource:          "172.30.30.58",
		DBInitial_Catalog: "golang_learning",
		DBUser_ID:         "sa",
		DBDriver:          "sqlserver",
		DBPassword:        "Mango@00$",
	}
}
