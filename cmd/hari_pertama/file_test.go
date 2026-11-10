package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInsertData(t *testing.T) {
	database, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, database.AutoMigrate(&User{}))

	user := User{
		Username: "andika",
		Nama:     "Andika",
		Password: "secret",
		Email:    "andika@example.com",
		Gender:   "l",
	}
	require.NoError(t, InsertData(database, user))

	var savedUser User
	require.NoError(t, database.First(&savedUser, "username = ?", user.Username).Error)
	require.Equal(t, user.Username, savedUser.Username)
	require.Equal(t, user.Nama, savedUser.Nama)
}
