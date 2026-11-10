package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func openDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := database.AutoMigrate(&User{}); err != nil {
		return nil, err
	}
	return database, nil
}

func InsertData(database *gorm.DB, usr User) error {
	return database.Create(&usr).Error
}
