package main

import (
	"fmt"
)

func main() {

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

		check := InsertData(fmt.Sprint(user) + "\n")
		if !check {
			fmt.Println("Gagal memasukkan data!")
		}

		fmt.Println("Lanjut (Y/n) : ")
		_, err = fmt.Scan(&pilihan)
		if err != nil {
			panic(err)
		}

	}

}
