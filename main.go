package main

import (
	"fmt"
	"os"
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
	filename := "data.txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

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

		file.WriteString(user.UserToString() + "\n")

		fmt.Println("Lanjut (Y/n) : ")
		_, err = fmt.Scan(&jawaban)
		if err != nil {
			panic(err)
		}
	}

}
