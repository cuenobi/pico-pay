package main

import (
	"log"
	"os"
	"strconv"

	"picopay/common/constant"
	database "picopay/common/databases"
	"picopay/common/models"
)

func main() {
	dbDialect := getEnv("DB_DIALECT", constant.POSTGRES)
	dbDSN := getEnv("DB_DSN", "postgres://admin:P@ssword1234@localhost:5432/pico-pay?sslmode=disable")
	maxOpen := getEnvAsInt("DB_MAX_OPEN", 10)
	maxIdle := getEnvAsInt("DB_MAX_IDLE", 5)
	maxLife := getEnvAsInt("DB_CONN_LIFETIME", 5)

	cfg := database.Config{
		Dialect:            dbDialect,
		DSN:                dbDSN,
		SetMaxOpenConns:    maxOpen,
		SetMaxIdleConns:    maxIdle,
		SetConnMaxLifetime: maxLife,
	}

	// Init DB
	if err := database.Init(cfg); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// Auto Migrate
	if err := database.DB.AutoMigrate(models.Models...); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("✅ Database migration complete!")
}

// --- Helpers ---

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func getEnvAsInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}
	return val
}
