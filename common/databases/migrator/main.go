package main

import (
	"fmt"
	"log"

	database "picopay/common/databases"
	"picopay/common/models"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	dbDialect := viper.GetString("DB.DIALECT")
	dbDSN := viper.GetString("DB.DSN")
	maxOpen := viper.GetInt("DB.MAX_OPEN")
	maxIdle := viper.GetInt("DB.MAX_IDLE")
	maxLife := viper.GetInt("DB.CONN_LIFETIME")

	cfg := database.Config{
		Dialect:            dbDialect,
		DSN:                dbDSN,
		SetMaxOpenConns:    maxOpen,
		SetMaxIdleConns:    maxIdle,
		SetConnMaxLifetime: maxLife,
	}

	fmt.Println(cfg)

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
