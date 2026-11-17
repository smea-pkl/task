package models

import "time"

// User adalah model tabel users pada database.
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Nama      string    `gorm:"not null"`
	Username  string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
}
