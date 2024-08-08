package config

import (
	"log"
	"os"

	"go-crud/utils"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DBType string
		Server *Server
		Db     *Db
	}

	Server struct {
		Port int64
	}

	Db struct {
		MongoDBURI  string
		PostgresDSN string
	}

	// Db struct {
	// 	Host     string
	// 	Port     int64
	// 	User     string
	// 	Password string
	// 	DBName   string
	// 	SSLMode  string
	// 	Timezone string
	// }
)

func LoadConfig(path string) *Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DBType: os.Getenv("DB_TYPE"),
		Server: &Server{
			Port: utils.ParseStringToInt(os.Getenv("SERVER_PORT")),
		},
		Db: &Db{
			MongoDBURI:  os.Getenv("DB_MONGODB_URI"),
			PostgresDSN: os.Getenv("DB_POSTGRESDSN"),
		},
	}
}
