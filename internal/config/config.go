package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string // dev, prod
	AppPort int

	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
}

func Load() (*Config, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	err = cfg.validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("load .env file: %w", err)
	}

	appPort, err := parsePort(os.Getenv("APP_PORT"))
	if err != nil {
		return nil, fmt.Errorf("parse APP_PORT: %w", err)
	}

	dbPort, err := parsePort(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("parse DB_PORT: %w", err)
	}

	cfg := &Config{
		AppEnv:  os.Getenv("APP_ENV"),
		AppPort: appPort,

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     dbPort,
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.AppEnv != "dev" && c.AppEnv != "prod" {
		return fmt.Errorf("APP_ENV must be 'dev' or 'prod'")
	}
	if c.DBHost == "" {
		return fmt.Errorf("DB_HOST is required")
	}
	if c.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	if c.DBUser == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if c.DBPassword == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	if c.DBSSLMode == "" {
		return fmt.Errorf("DB_SSLMODE is required")
	}

	return nil
}

func parsePort(portStr string) (int, error) {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, fmt.Errorf("parse port must be a valid integer: %w", err)
	}
	if port <= 0 || port > 65535 {
		return 0, fmt.Errorf("port must be between 1 and 65535")
	}
	return port, nil
}