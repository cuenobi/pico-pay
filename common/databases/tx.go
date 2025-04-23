package database

import "gorm.io/gorm"

func WithTx(fn func(tx *gorm.DB) error) error {
    return DB.Transaction(fn)
}