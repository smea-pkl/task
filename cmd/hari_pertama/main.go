package main

import (
	"fmt"
)

func main() {
	database, err := openDatabase()
	if err != nil {
		panic(err)
	}

	// Migrasi data dari file txt ke SQLite (database) sebelum mulai menerima input.
	db()

	pilihan := "Y"
	for pilihan == "Y" || pilihan == "y" {
		user, err := NewUser()
		if err != nil {
			panic(err)
		}

		allowed, err := FilterUser(user)
		if err != nil {
			panic(err)
		}

		if !allowed {
			fmt.Println("Kamu tidak boleh daftar")
		} else {
			if err := InsertData(database, user); err != nil {
				fmt.Println("Gagal memasukkan data:", err)
			}
		}

		fmt.Print("Lanjut (Y/n) : ")
		_, err = fmt.Scan(&pilihan)
		if err != nil {
			panic(err)
		}
	}
}
