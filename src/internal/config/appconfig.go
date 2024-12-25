package config

import (
	"os"

	"gorm.io/gorm"
)

type dbConfig struct {
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBNAME     string
	DBPORT     string
}

type serverConfig struct {
	BACKEND_URL string
}

type googleConfig struct {
	GOOGLE_CLIENT_ID  string
	GOOGLE_SECRET_KEY string
}

type jwtConfig struct {
	JWT_SECRET_KEY          string
	JWT_EXPIRATION_IN_HOURS string
}

type AppConfig struct {
	Server       serverConfig
	DBConfig     dbConfig
	GoogleConfig googleConfig
	JWTConfig    jwtConfig
	DB           *gorm.DB
}

func NewConfig() *AppConfig {
	appConfig := AppConfig{

		Server: serverConfig{
			BACKEND_URL: getEnv("BACKEND_URL", ""),
		},

		DBConfig: dbConfig{
			DBUSER:     getEnv("DB_USERNAME", ""),
			DBPASSWORD: getEnv("DB_PASSWORD", ""),
			DBNAME:     getEnv("DB_NAME", ""),
			DBHOST:     getEnv("DB_HOST", ""),
			DBPORT:     getEnv("DB_PORT", ""),
		},

		GoogleConfig: googleConfig{
			GOOGLE_CLIENT_ID:  getEnv("GOOGLE_CLIENT_ID", ""),
			GOOGLE_SECRET_KEY: getEnv("GOOGLE_SECRET_KEY", ""),
		},

		JWTConfig: jwtConfig{
			JWT_SECRET_KEY:          getEnv("JWT_SECRET_KEY", ""),
			JWT_EXPIRATION_IN_HOURS: getEnv("JWT_EXPIRATION_IN_HOURS", ""),
		},
	}
	return &appConfig
}

func getEnv(key string, defaultval string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultval
}
