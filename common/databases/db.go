package database

import (
	"fmt"
	"log"
	"time"

	"picopay/common/constant"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
	Dialect            string
	DSN                string
	SetMaxOpenConns    int
	SetMaxIdleConns    int
	SetConnMaxLifetime int
}

func Init(cfg Config) error {
	var dialector gorm.Dialector
	switch cfg.Dialect {
	case constant.POSTGRES:
		dialector = postgres.Open(cfg.DSN)
	case constant.MYSQL:
		dialector = mysql.Open(cfg.DSN)
	default:
		return ErrUnsupportedDialect
	}

	var err error
	DB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(cfg.SetMaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.SetMaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(cfg.SetConnMaxLifetime))

	log.Println("Connected to DB using GORM")
	return nil
}

var ErrUnsupportedDialect = fmt.Errorf("unsupported database dialect")
