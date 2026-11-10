package main

import (
	"fmt"
	"os"
)

func InsertData(usr string) bool {
	filename := "user_data.txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return false
	}
	defer file.Close()

	file.WriteString(usr)
	
	return true
}
