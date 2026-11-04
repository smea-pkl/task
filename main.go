package main

import (
	"fmt"
)

func main() {
	// var nama string

	// fmt.Print("Masukkan nama anda : ")
	// _, err := fmt.Scan(&nama)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Halo", nama)

	jawaban := "Y"

Parent:
	for jawaban == "Y" || jawaban == "y" {

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

		check := InsertData(user.UserToString() + "\n")
		if !check {
			fmt.Println("Gagal memasukkan data!")
		}

		fmt.Println("Lanjut (Y/n) : ")
		_, err = fmt.Scan(&jawaban)
		if err != nil {
			panic(err)
		}
	}

}
