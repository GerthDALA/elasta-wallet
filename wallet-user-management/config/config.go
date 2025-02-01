package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	RedisHost  string
	RedisPort  int
	JWTSecret  string
	AppPort    string
}

func LoadConfig() (*Config, error) {
	
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}

	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_PORT: %v", err)
	}

	config := &Config{
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     port,
		RedisHost:  os.Getenv("REDIS_HOST"),
		RedisPort:  redisPort,
		JWTSecret:  os.Getenv("JWT_SECRET"),
		AppPort:    os.Getenv("APP_PORT"),
	}

	return config, nil

}
