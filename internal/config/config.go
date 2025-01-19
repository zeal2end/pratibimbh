package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the main configuration for the application
type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	DatabaseURL string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Load returns configuration loaded from environment variables
func Load() *Config {
	// Load the .env file only if it exists (for local development)
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
		} else {
			fmt.Println("Loaded .env file successfully")
		}
	}

	return &Config{
		Server:      loadServerConfig(),
		Database:    loadDatabaseConfig(),
		DatabaseURL: GetDSN(),
	}
}

func loadServerConfig() ServerConfig {
	return ServerConfig{
		Port: getEnv("SERVER_PORT", "8080"),
	}
}

func loadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "19852"),
		User:     getEnv("DB_USER", "bloguser"),
		Password: getEnv("DB_PASSWORD", "blogpass"),
		DBName:   getEnv("DB_NAME", "blogdb"),
		SSLMode:  getEnv("DB_SSLMODE", "enable"),
	}
}

// Helper function to get environment variables with default values
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetDSN returns database connection string
func GetDSN() string {
	// Always try to get DB_URL from environment variable
	return getEnv("DB_URL", "no url")
}
