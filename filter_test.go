package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInput(t *testing.T) {
	t.Run("User input sutarman", func(t *testing.T) {
		usr := User{
			name:     "Sutarman",
			username: "Sutarman",
			email:    "sutarman@gmail.com",
			password: "sutarman",
			gender:   "l",
		}

		check, err := FilterUser(usr)
		assert.False(t, check)
		assert.Nil(t, err)
	})

	t.Run("User input andika", func(t *testing.T) {
		usr := User{
			name:     "Andika",
			username: "Andika",
			email:    "andika@gmail.com",
			password: "andika",
			gender:   "l",
		}

		check, err := FilterUser(usr)
		assert.False(t, check)
		assert.Nil(t, err)
	})

	t.Run("User input fitri dan gender p", func(t *testing.T) {
		usr := User{
			name:     "Fitri",
			username: "Fitri",
			email:    "fitri@gmail.com",
			password: "fitri",
			gender:   "p",
		}

		check, err := FilterUser(usr)
		assert.False(t, check)
		assert.Nil(t, err)
	})

	t.Run("User input fitri dan gender l", func(t *testing.T) {
		usr := User{
			name:     "Fitri",
			username: "Fitri",
			email:    "fitri@gmail.com",
			password: "fitri",
			gender:   "l",
		}

		check, err := FilterUser(usr)
		assert.True(t, check)
		assert.Nil(t, err)
	})

	t.Run("User input gender selain p, P, l, L", func(t *testing.T) {
		usr := User{
			name:     "Rizki",
			username: "rizki",
			email:    "rizki@gmail.com",
			password: "rizki",
			gender:   "T",
		}

		check, err := FilterUser(usr)
		assert.False(t, check)
		assert.Nil(t, err)
	})
}
