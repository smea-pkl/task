package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"pkl/backend/models"
)

// Init membuka koneksi SQLite dan memastikan tabel sudah dimigrasi.
func Init(file string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	return db, nil
}
