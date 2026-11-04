package main

import "fmt"

type User struct {
	name     string
	username string
	email    string
	password string
	gender   string
}

func NewUser() (User, error) {
	var user User

	fmt.Print("Masukkan Nama : ")
	_, err := fmt.Scanln(&user.name)

	if err != nil {
		return user, err
	}

	fmt.Print("Masukkan Username : ")
	_, err = fmt.Scanln(&user.username)

	if err != nil {
		return user, err
	}

	fmt.Print("Masukkan email : ")
	_, err = fmt.Scanln(&user.email)
	if err != nil {
		return user, err
	}

	fmt.Print("Masukkan password : ")
	_, err = fmt.Scanln(&user.password)

	if err != nil {
		return user, err
	}

	fmt.Print("Masukkan gender : ")
	_, err = fmt.Scanln(&user.gender)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (usr User) UserToString() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s", usr.name, usr.username, usr.email, usr.password, usr.gender)
}
