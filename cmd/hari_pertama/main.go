package main

import (
	"fmt"
)

func main() {
	database, err := openDatabase()
	if err != nil {
		panic(err)
	}

	pilihan := "Y"
Parent:
	for pilihan == "Y" || pilihan == "y" {

		var user, err = NewUser()
		if err != nil {
			panic(err)
		}

		b, err := FilterUser(user)
		if err != nil {
			panic(err)
		}
		if b == false {
			fmt.Println("Kamu tidak boleh daftar")
			continue Parent
		}

		if err := InsertData(database, user); err != nil {
			fmt.Println("Gagal memasukkan data:", err)
		}

		fmt.Println("Lanjut (Y/n) : ")
		_, err = fmt.Scan(&pilihan)
		if err != nil {
			panic(err)
		}

	}

}
