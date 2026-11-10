package main

import "fmt"

type User struct {
	Name     string
	Username string
}

func main() {
	users := []User{
		{Name: "anita", Username: "sdasd"},
		{Name: "dita", Username: "asd"},
		{Name: "anita", Username: "3"},
		{Name: "anita", Username: "1"},
		{Name: "dita", Username: "4"},
		{Name: "j", Username: "fgh"},
	}

	uniqueUsers := make([]User, 0, len(users))
	seen := make(map[string]bool)

	for _, user := range users {
		key := user.Name + ":" + user.Username
		if seen[key] {
			continue
		}
		seen[key] = true
		uniqueUsers = append(uniqueUsers, user)
	}

	fmt.Println(uniqueUsers)
}

// contoh ruwet
// func main() {

// 	users := []string{
// 		"anita",
// 		"fitri",
// 		"asdasd",
// 		"none",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"anita",
// 		"fitri",
// 		"asdasd",
// 		"none",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"fitri",
// 		"asdasd",
// 		"james",
// 		"anita",
// 		"fitri",
// 		"asdasd",
// 		"none",
// 		"fitri",
// 		"asdasd",
// 	}

// 	user_unik := []string{}

// 	// 1. ambil satu persatu users.
// 	for _, item := range users {

// 		var duplikat bool = false
// 		for _, unik := range user_unik {
// 			// 2. check per item di user unik
// 			if item == unik {
// 				duplikat = true
// 			}
// 		}

// 		// 3. kalau ada dibuang
// 		// 4. kalau tidak ada, dimasukkan.
// 		if duplikat {
// 			continue
// 		} else {
// 			user_unik = append(user_unik, item)
// 		}

// 	}

// 	// tampilkan hasilnya
// 	fmt.Println(user_unik)

// }
