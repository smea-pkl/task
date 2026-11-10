package main

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string 
	Nama     string 
	Password string
	Email    string
	Gender   string
}

func NewUser() (User, error) {
	var user User
	fmt.Print("Masukkan username : ")  //1.tanya username
	_, err := fmt.Scan(&user.Username) //2.input username
	if err != nil {
		panic(err)
	}

	fmt.Print("Masukkan nama : ") //3.tanya nama
	_, err = fmt.Scan(&user.Nama) //4.input nama
	if err != nil {
		panic(err)
	}

	fmt.Print("Masukkan password : ") //5.tanya password
	_, err = fmt.Scan(&user.Password) //6.input password
	if err != nil {
		panic(err)
	}

	fmt.Print("Masukkan email : ") //7.tanya email
	_, err = fmt.Scan(&user.Email) //8.input email
	if err != nil {
		panic(err)
	}

	fmt.Print("Masukkan gender (l/p) : ") //9.tanya gender
	_, err = fmt.Scan(&user.Gender)       //10.input gender
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (usr User) UserToString() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s", usr.Nama, usr.Username, usr.Email, usr.Password, usr.Gender)
}
